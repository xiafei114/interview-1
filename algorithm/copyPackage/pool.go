package copyPackage

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)


var allPool []*Pool
var allPoolMu sync.Mutex

const dequeueBits = 32

type Pool struct {
	local unsafe.Pointer
	localSize uintptr

	victim unsafe.Pointer
	victimSize uintptr

	New func() interface{}

	seed uint32
}

type poolLocal struct {
	poolLocalInternal

	// 填充位, 防止cache missing
	pad [128 - unsafe.Sizeof(poolLocalInternal{}) % 128]byte
}

type poolLocalInternal struct {
	private interface{}
	shared poolChain
}

type poolChain struct {
	head *poolChainElt

	tail *poolChainElt
}

type poolChainElt struct {
	poolDequeue

	next, prev *poolChainElt
}

type poolDequeue struct {
	headTail uint64

	vals []interface{}
}

func (p *Pool) Get() interface{}{
	// 获取pid和pid对应的local
	l, pid := p.pin()
	x := l.private
	l.private = nil

	// private 没取到
	if x == nil {
		// 从shared上面取
		x, _ = l.shared.popHead()
		if x == nil {
			// 从其他pid下面的shared偷一个
			x = p.getSlow(pid)
		}
	}

	// new函数不为空, 实例化一个新的
	if x == nil && p.New != nil {
		x = p.New()
	}

	return x
}

func (p *Pool) Put(x interface{}) {
	if x == nil {
		return
	}

	// 获取当前pid对应的poolLocal
	l, _ := p.pin()
	if l.private == nil {
		l.private = x
		x = nil
	}

	if x != nil {
		l.shared.pushHead(x)
	}
}

func indexLocal(l unsafe.Pointer, i int) *poolLocal {
	// 获取对应pid的poolLocal
	return (*poolLocal)(unsafe.Pointer(uintptr(l) + uintptr(i) * unsafe.Sizeof(poolLocal{})))
}

func (p *Pool) pin() (*poolLocal, int) {
	size := atomic.LoadUintptr(&p.localSize)
	l := p.local

	if atomic.CompareAndSwapUint32(&p.seed, 0, 1) {
		rand.Seed(time.Now().Unix())
	}

	pid := p.getPid()
	if uintptr(pid) < size {
		return indexLocal(l, pid), pid
	}

	return p.pinSlow()
}

func (p *Pool) pinSlow() (*poolLocal, int){
	allPoolMu.Lock()
	defer allPoolMu.Unlock()

	pid := p.getPid()
	s := p.localSize
	l := p.local
	if uintptr(pid) < s {
		return indexLocal(l, pid), pid
	}

	// 如果当前local是空的, 要把这个pool加入到allPool
	if p.local == nil {
		allPool = append(allPool, p)
	}

	size := runtime.GOMAXPROCS(0)
	local := make([]poolLocal, size)
	atomic.StorePointer(&p.local, unsafe.Pointer(&local))
	atomic.StoreUintptr(&p.localSize, uintptr(size))

	return &local[pid], pid
}

func (p *Pool) getPid() int {
	procNum := runtime.GOMAXPROCS(0)
	return rand.Intn(procNum)
}



// get slow在调用时， 已经进行了goroutine和p的绑定， 因此不需要加锁
func (p *Pool) getSlow(pid int) interface{} {
	// 去其他p的shared偷
	size := atomic.LoadUintptr(&p.localSize)
	locals := p.local
	for i:=0;i<int(size);i++ {
		// (pid+i+1)%int(size) 可以计出除了本身之外所有的0～size-1
		l := indexLocal(locals, (pid+i+1) % int(size))
		if x, _ := l.shared.popTail(); x != nil {
			return x
		}
	}

	// shared拿不到, 从本地的victim取
	size = atomic.LoadUintptr(&p.victimSize)
	if uintptr(pid) >= size {
		return nil
	}

	locals = p.victim
	// 取自己的victim
	l := indexLocal(locals, pid)
	if x := l.private; x != nil {
		l.private = nil
		return x
	}

	// 从其他人的victim里面偷
	for i:=0;i<int(size);i++ {
		l := indexLocal(locals, (pid+i+1)%int(size))
		if x, _ := l.shared.popTail(); x != nil {
			return x
		}
	}

	// 如果都没获取到, 清空victim, 在下次GC前就不用在获取了
	atomic.StoreUintptr(&p.victimSize, 0)
	return nil
}
func (p *poolChain) popTail() (interface{}, bool) {
	d := loadPoolChainElt(&p.tail)
	if d == nil {
		return nil, false
	}
	for {
		// 尾部找到了， 直接返回
		if val, ok := d.popTail(); ok {
			return val, ok
		}

		d2 := loadPoolChainElt(&d.next)
		if d2 == nil {
			return nil, false
		}

		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&p.tail)), unsafe.Pointer(d), unsafe.Pointer(&d2)) {
			storePoolChainElt(&d2.prev, nil)
		}

	}
}


func (pc *poolChain) popHead() (interface{}, bool) {
	d := pc.head
	for d != nil {
		if val, ok := d.popHead(); ok {
			return val, ok
		}

		d = loadPoolChainElt(&d.prev)
	}

	return nil, false
}

func loadPoolChainElt(elt **poolChainElt) *poolChainElt {
	return (*poolChainElt)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(elt))))
}
func storePoolChainElt(elt **poolChainElt, v *poolChainElt) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(elt)), unsafe.Pointer(v))
}

func (pce *poolChainElt) popHead() (interface{}, bool) {
	var slot interface{}
	for {
		// 解析出 head, tail 指针
		// 判断当前这个环是否满了
		ptrs := atomic.LoadUint64(&pce.headTail)
		head, tail := pce.unpack(ptrs)
		if tail == head {
			return nil, false
		}

		head--
		ptrs2 := pce.pack(head, tail)
		if atomic.CompareAndSwapUint64(&pce.headTail, ptrs, ptrs2) {
			slot = &pce.vals[head&uint32(len(pce.vals) - 1)]
			break
		}
	}

	val := *(*interface{})(unsafe.Pointer(&slot))

	return val, true
}

func (pd *poolDequeue) pushHead(x interface{}) bool {
	ptrs := atomic.LoadUint64(&pd.headTail)
	head, tail := pd.unpack(ptrs)
	if (tail + uint32(len(pd.vals)) & (1<<dequeueBits - 1)) == head {
		return false
	}


	// 因为head是存储在高32位的, 所以低32位+1向33位进位
	atomic.AddUint64(&pd.headTail, 1<<dequeueBits)
	return true
}

func (pce *poolDequeue) pack(head, tail uint32) uint64 {
	const mask = 1 << dequeueBits - 1 // 31位
	return (uint64(head) << dequeueBits) | uint64(tail&mask)
}

func (pce *poolDequeue) unpack(ptrs uint64) (head, tail uint32) {
	const mask = 1 << dequeueBits - 1
	head = uint32((head << dequeueBits)&mask)
	tail = uint32(tail&mask)

	return
}


func (pc *poolChain) pushHead(x interface{}) {
	d := pc.head
	// 空值, 去初始化
	if d == nil {
		const initSize = 8
		d = new(poolChainElt)
		d.vals = make([]interface{}, initSize)
		// ?? 这里为啥要这么存
		pc.head = d
		storePoolChainElt(&pc.tail, d)
	}

	if d.pushHead(x) {
		return
	}

	// push不进去是上一个ele满了, 创建一个新的
	newSize := len(d.vals) * 2
	if newSize > (1 << dequeueBits) / 4 {
		newSize = (1<<dequeueBits)/4
	}

	d2 := &poolChainElt{prev: d}
	d2.vals = make([]interface{}, newSize)
	pc.head = d2
	storePoolChainElt(&d.next, d2)
	d2.pushHead(x)
}


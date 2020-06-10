1. 简述如何设计一个秒杀系统
2. 线程, 协程, 进程的区别
3. 如何设计实现一个channel
4. 10个袋子各装有一颗宝石, 9个是真的,1个是假的, 怎快最快找出假的
5. fabonacci的递归算法O(2^n)以及非递归算法O(n)
6. name 是索引字段, select id, name 以及 select id, name, age 的区别
7. 再有缓存的情况下, 更新数据时, 怎么保证数据的一致性
8. qps = 10的情况下, 是先删除缓存在update DB还是 先update DB在更新缓存
9. linux怎么查询进程打开了多少个文件句柄
	`ulimit -u` 查看系统默认的最大文件句柄数
	`lsof -n | awk '{print $2}' | uniq -c | wc -l` 查看当前打开的文件句柄数
10. 给定非负整型数组arr和整数limit，两次从arr中随机抽取元素（可能抽到一个同元素），获得整数x和y，计算和s=x+y
, 求所有不超过limit的s值中最大数
11. 假设我们有8种不同的钱币面值{1，2，5，10，20，50，100，200}，用这些钱币组合成一个给定的数值n, 求出有多少种组合
12. 进程
	```
		struct task_struct {
			long state; # 进程状态
			struct mm_struct *mm; 虚拟内存
			pid_t pid;
			struct task_struct *parent; 父进程指针
			struct list_head children; 子进程列表
			struct fs_struct *fs; 存放文件系统信息的指针
			struct file_struct *files; 数组， 包含该进程打开的文件列表
		}
	```
13. 文件描述符: 进程的files字段是一个数组，一个进程一般会从files[0](标准输入)读取文件, 向files[1](标准输出)或者files[2](标准错误输出)写入数据,当该进程希望从新的文件中读取数据时, 内核把文件打开，放在files[3]的位置，同理，输入输出的重定向只需要更改files[0~2]的值就可以了


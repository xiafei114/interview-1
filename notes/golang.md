1. golang 的runtime机制(待补充完整)
runtime负责和底层操作系统交互

2. for...range 的坑, 值的拷贝,无法改变原始数据
```
package main
import "fmt"

func main() {
    pase_student()
}

type student struct {
    Name string
    Age  int 
}

func pase_student() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }   
    // stu是临时变量,地址值不变, 所以&stu会把相同的地址值一直带到最后一次循环
    for _, stu := range stus {

        fmt.Printf("%+v\n", &stu)
        fmt.Printf("%+v\n", stu)
		tmp := stu
        m[stu.Name] = &tmp
    }

    for k,v:=range m{
        println(k,"=>",v.Name)
    }

}

```

3. 结构体进行比较的时候, 不但与属性类型个数有关,还与属性顺序相关, reflect.DeepEqual可以对map, slice等结构进行比较



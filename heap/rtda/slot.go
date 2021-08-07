package rtda

import "JVM-Go/heap/rtda/heap"

//Slot 局部变量表数组实现的基本单位，既可以存放一个int值也可以存放一个引用值
type Slot struct {
	num int32        //存放整数
	ref *heap.Object //存放引用
}

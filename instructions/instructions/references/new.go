package references

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
)

//NEW 创建实例对象指令
type NEW struct {
	base.Index16Instruction //new指令的操作数是一个uint16索引，来自字节码
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.GetMethod().GetClass().GetConstantPool()
	classRef := cp.GetConstant(self.Index).(heap.ClassRef) //通过索引，可以从当前类的运行时常量池中找到一个类符号引用
	class := classRef.ResolvedClass()                      //解析这个类符号引用
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError") //接口和抽象类都不能实例化
	}
	ref := class.NewObject()             //拿到类数据，然后创建对象
	frame.GetOperandStack().PushRef(ref) //把对象引用推入栈顶
}

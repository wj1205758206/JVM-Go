package references

import (
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
	"JVM-Go/heap/instructions/base"
)

//INSTANCE_OF 判断对象是否是某个类的实例（或者对象的类是 否实现了某个接口）并把结果推入操作数栈
type INSTANCE_OF struct {
	//第一个操作数是uint16索引， 从方法的字节码中获取，通过这个索引可以从当前类的运行时常量 池中找到一个类符号引用
	//第二个操作数是对象引用，从操作数栈中弹出
	base.Index16Instruction
}

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	cp := frame.GetMethod().GetClass().GetConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	stack := frame.GetOperandStack()
	ref := stack.PopRef() //弹出对象引用，如果是null，则把0推入操作数栈
	if ref == nil {
		stack.PushInt(0)
		return
	}
	if ref.IsInstanceOf(class) { //如果对象引用不是null，则解析类符号引用得到的class，判断对象是否是类的实例，然后把判断结果推入操作数栈
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}

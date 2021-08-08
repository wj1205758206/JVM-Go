package references

import (
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
	"JVM-Go/instructions/instructions/base"
)

//CHECK_CAST 不改变操作数栈,如果判断失败，直接抛出异常
type CHECK_CAST struct {
	base.Index16Instruction
}

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref) //先从操作数栈中弹出对象引用，再推回去，这样就不会改变操作数栈的状态
	if ref == nil {
		return //如果引用是null，则指令执行结束。也就是说，null 引用可以转换成任何类型
	}

	cp := frame.GetMethod().GetClass().GetConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) { //解析类符号引用，判断对象是否是类的实例。如果是的话，指令执行结束，否则抛出异常
		panic("java.lang.ClassCastException")
	}
}

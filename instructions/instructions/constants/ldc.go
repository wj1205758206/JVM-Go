package constants

import (
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
	"JVM-Go/instructions/instructions/base"
)

//ldc系列指令从运行时常量池中加载常量值，并把它推入操作数栈

//LDC 用于加载int、float和字符串常量
type LDC struct {
	base.Index8Instruction
}

//LDC_W java.lang.Class实例或者MethodType和MethodHandle实例
type LDC_W struct {
	base.Index16Instruction
}

//LDC2_W 指令用于加载long和double常量
type LDC2_W struct {
	base.Index16Instruction
}

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.GetOperandStack()
	cp := frame.GetMethod().GetClass().GetConstantPool()
	constant := cp.GetConstant(index)
	switch constant.(type) { //先从当前类的运行时常量池中取出常量，按照类型推入操作数栈
	case int32:
		stack.PushInt(constant.(int32))
	case float32:
		stack.PushFloat(constant.(float32))
	case string:
		//TODO
	case *heap.ClassRef:
		//TODO
	default:
		panic("TODO: ldc!")
	}
}

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	cp := frame.GetMethod().GetClass().GetConstantPool()
	constant := cp.GetConstant(self.Index)
	switch constant.(type) {
	case int64:
		stack.PushLong(constant.(int64))
	case float64:
		stack.PushDouble(constant.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

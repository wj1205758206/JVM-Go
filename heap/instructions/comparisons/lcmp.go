package comparisons

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//LCMP 比较两个long类型变量的大小
type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}

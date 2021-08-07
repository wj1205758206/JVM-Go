package conversions

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//F2D float类型转成double
type F2D struct {
	base.NoOperandsInstruction
}

//F2L float类型转成long
type F2L struct {
	base.NoOperandsInstruction
}

//F2I float类型转成int
type F2I struct {
	base.NoOperandsInstruction
}

func (self *F2D) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	fVal := stack.PopFloat()
	dVal := float64(fVal)
	stack.PushDouble(dVal)
}

func (self *F2L) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	fVal := stack.PopFloat()
	lVal := int64(fVal)
	stack.PushLong(lVal)
}

func (self *F2I) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	fVal := stack.PopFloat()
	iVal := int32(fVal)
	stack.PushInt(iVal)
}

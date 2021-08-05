package conversions

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//D2F double类型转成float
type D2F struct {
	base.NoOperandsInstruction
}

//D2L double类型转成long
type D2L struct {
	base.NoOperandsInstruction
}

//D2I double类型转成int
type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	dVal := stack.PopDouble()
	fVal := float32(dVal)
	stack.PushFloat(fVal)
}

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	dVal := stack.PopDouble()
	lVal := int64(dVal)
	stack.PushLong(lVal)
}

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	dVal := stack.PopDouble()
	iVal := int32(dVal)
	stack.PushInt(iVal)
}

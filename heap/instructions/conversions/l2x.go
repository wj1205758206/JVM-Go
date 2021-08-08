package conversions

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//L2I long类型转int
type L2I struct {
	base.NoOperandsInstruction
}

//L2F long类型转float
type L2F struct {
	base.NoOperandsInstruction
}

//L2D long类型转double
type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	lVal := stack.PopLong()
	iVal := int32(lVal)
	stack.PushInt(iVal)
}

func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	lVal := stack.PopLong()
	fVal := float32(lVal)
	stack.PushFloat(fVal)
}

func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	lVal := stack.PopLong()
	dVal := float64(lVal)
	stack.PushDouble(dVal)
}

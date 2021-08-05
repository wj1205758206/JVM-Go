package conversions

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//I2S int类型转short
type I2S struct {
	base.NoOperandsInstruction
}

//I2C int类型转char
type I2C struct {
	base.NoOperandsInstruction
}

//I2B int类型转byte
type I2B struct {
	base.NoOperandsInstruction
}

//I2L int类型转long
type I2L struct {
	base.NoOperandsInstruction
}

//I2F int类型转float
type I2F struct {
	base.NoOperandsInstruction
}

//I2D int类型转double
type I2D struct {
	base.NoOperandsInstruction
}

func (self *I2S) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	iVal := stack.PopInt()
	sVal := int32(int16(iVal))
	stack.PushInt(sVal)
}

func (self *I2C) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	iVal := stack.PopInt()
	cVal := int32(int16(iVal))
	stack.PushInt(cVal)
}

func (self *I2B) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	iVal := stack.PopInt()
	bVal := int32(int8(iVal))
	stack.PushInt(bVal)
}

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	iVal := stack.PopInt()
	lVal := int64(iVal)
	stack.PushLong(lVal)
}

func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	iVal := stack.PopInt()
	fVal := float32(iVal)
	stack.PushFloat(fVal)
}

func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	iVal := stack.PopInt()
	dVal := float64(iVal)
	stack.PushDouble(dVal)
}

package constants

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//BIPUSH
type BIPUSH struct {
	val int8
}

//SIPUSH
type SIPUSH struct {
	val int16
}

//从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.GetOperandStack().PushInt(i)
}

//从操作数中获取一个short型整数，扩展成int型，然后推入栈顶
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.GetOperandStack().PushInt(i)
}

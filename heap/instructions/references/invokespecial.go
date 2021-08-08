package references

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PopRef()
}

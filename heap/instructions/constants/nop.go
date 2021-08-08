package constants

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//NOP 空指令不执行任何操作
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}

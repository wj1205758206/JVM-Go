package control

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//GOTO 进行无条件跳转
type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}

package control

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//LOOKUP_SWITCH 多条件分支跳转，case值是不连续的
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

//Execute 从操作数栈中弹出一个int变量，然后用它查找 matchOffsets，看是否能找到匹配的key。如果能，则按照value给出的 偏移量跳转，否则按照defaultOffset跳转
func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	key := stack.PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}

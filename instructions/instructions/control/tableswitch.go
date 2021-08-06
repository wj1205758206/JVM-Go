package control

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//TABLE_SWITCH 多条件分支跳转，case值必须是连续的
type TABLE_SWITCH struct {
	defaultOffset int32 //应默认情况下执行跳转所需的字节码偏移量
	low           int32 //low和high记录case的取值范围
	high          int32
	jumpOffsets   []int32 //jumpOffsets是一个索引表，里面存放high-low+1个int值，对应各种case情况下，执行跳转所需的字节码偏移量
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding() //tableswitch指令操作码的后面有0~3字节的padding，以保证defaultOffset在字节码中的地址是4的倍数
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetCount)
}
//Execute 从操作数栈中弹出一个int变量，然后看它是 否在low和high给定的范围之内。如果在，则从jumpOffsets表中查出 偏移量进行跳转，否则按照defaultOffset跳转
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	index := stack.PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}

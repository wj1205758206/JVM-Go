package extended

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/instructions/loads"
	"JVM-Go/heap/instructions/math"
	"JVM-Go/heap/instructions/stores"
	"JVM-Go/heap/rtda"
)

//WIDE wide指令改变其他指令的行为
type WIDE struct {
	modifiedInstruction base.Instruction //存放被改变的指令
}

//FetchOperands 从字节码中读取一字节的操作码，然 后创建子指令实例，最后读取子指令的操作数
func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: //iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16()) //索引原来只有1字节，现在扩展成2字节
		self.modifiedInstruction = inst
	case 0x16: //lload
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x17: //fload
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x18: //dload
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x19: //aload
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x36: //istore
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x37: //lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x38: //fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x39: //dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x3a: //astore
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x84: //iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0xa9: //ret

	}
}

//Execute wide指令只是增加了索引宽度，并不改变子指令操作
func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}

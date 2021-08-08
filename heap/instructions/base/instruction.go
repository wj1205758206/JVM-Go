package base

import "JVM-Go/heap/rtda"

//Instruction 把指令抽象成接口，解码和执行逻辑写在具体的指令实现中
type Instruction interface {
	FetchOperands(reader *BytecodeReader) //从字节码中提取操作数
	Execute(frame *rtda.Frame)            //执行指令逻辑
}

//NoOperandsInstruction 无操作数指令
type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

//BranchInstruction 跳转指令
type BranchInstruction struct {
	Offset int //存放跳转偏移量
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

//Index8Instruction 存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出
type Index8Instruction struct {
	Index uint //表示局部变量表索引
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

//Index16Instruction 有一些指令需要访问运行时常量池，常量池索引由两字节操作数给出
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}

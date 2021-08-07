package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

//ConstantFieldrefInfo 结构体定义，嵌套ConstantMemberrefInfo结构体
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

//ConstantMethodrefInfo 结构体定义，嵌套ConstantMemberrefInfo结构体
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

//ConstantInterfaceMethodrefInfo 结构体定义，嵌套ConstantMemberrefInfo结构体
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

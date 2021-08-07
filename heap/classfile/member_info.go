package classfile

//MemberInfo 字段表和方法表的结构非常类似，统一抽取为MemberInfo结构体
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

//readMembers 读取字段表数据或者方法表数据
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

//readMember 读取字段或者方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

//GetAccessFlags 获取访问标识
func (self *MemberInfo) GetAccessFlags() uint16 {
	return self.accessFlags
}

//GetName 从常量池查找字段或方法名
func (self *MemberInfo) GetName() string {
	return self.cp.getUtf8(self.nameIndex)
}

//GetDescriptor 从常量池查找字段或方法描述符
func (self *MemberInfo) GetDescriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) GetCodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) GetConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

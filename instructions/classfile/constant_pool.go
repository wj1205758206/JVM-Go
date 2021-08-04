package classfile

//ConstantPool 自定义ConstantPool类型
type ConstantPool []ConstantInfo

//readConstantPool 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16()) //常量池计数器
	cp := make([]ConstantInfo, cpCount) //常量池表
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo: //CONSTANT_Long_info和CONSTANT_Double_info在常量池中各占两个位置
			i++
		}
	}
	return cp
}

//getConstantInfo 按索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}

//getNameAndType 从常量池查找字段或方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	nameAndTypeInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(nameAndTypeInfo.nameIndex)
	_type := self.getUtf8(nameAndTypeInfo.descriptorIndex)
	return name, _type
}

//getClassName 从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classNameInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classNameInfo.nameIndex)
}

//getUtf8 从常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

//ConstantInfo 该接口用来表示常量池表项信息
type ConstantInfo interface {
	//readInfo 该方法读取常量池表项信息，需要由具体的常量结构体实现
	readInfo(reader *ClassReader)
}

//readConstantInfo 该方法先读出tag值，然后调用newConstantInfo函数创建具体的常量，最后调用常量的readInfo方法读取常量信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

//newConstantInfo 根据tag值创建具体的常量
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}

package classfile

import "fmt"

//ClassFile Class文件结构体
type ClassFile struct {
	magic        uint32           //魔数
	minorVersion uint16           //副版本号
	majorVersion uint16           //主版本号
	constantPool ConstantPool //常量池表
	accessFlags  uint16           //访问标识
	thisClass    uint16           //类索引
	superClass   uint16           //父类索引
	interfaces   []uint16         //接口索引集合
	fileds       []*MemberInfo    //字段表
	methods      []*MemberInfo    //方法表
	attributes   []AttributeInfo  //属性表
}

//Parse 将[]byte二进制数据解析后才能ClassFile结构体
func Parse(classdata []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classdata}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

//read 该方法一次调用其它方法解析class文件
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16Table()
	self.fileds = readMembers(reader, self.constantPool)
	self.methods = readAttributes(reader, self.constantPool)
}

//readAndCheckMagic 读取魔数并检查
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

//readAndCheckVersion 读取并检查版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

//GetMinorVersion 获取副版本号
func (self *ClassFile) GetMinorVersion() uint16 {
	return self.minorVersion
}

//GetMajorVersion 获取主版本号
func (self *ClassFile) GetMajorVersion() uint16 {
	return self.majorVersion
}

//GetConstantPool 获取常量池
func (self *ClassFile) GetConstantPool() ConstantPool {
	return self.constantPool
}

//GetAccessFlags 获取访问标识
func (self *ClassFile) GetAccessFlags() uint16 {
	return self.accessFlags
}

//GetFields 获取字段表
func (self *ClassFile) GetFields() []*MemberInfo {
	return self.fileds
}

//GetMethods 获取方法表
func (self *ClassFile) GetMethods() []*MemberInfo {
	return self.methods
}

//GetClassName 从常量池中获取类名
func (self *ClassFile) GetClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

//GetSuperClassName 从常量池中获取父类名
func (self *ClassFile) GetSuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //Object没有父类
}

//GetInterfaceNames 从常量池查找接口名
func (self *ClassFile) GetInterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

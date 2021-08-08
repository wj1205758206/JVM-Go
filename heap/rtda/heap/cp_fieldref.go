package heap

import "JVM-Go/heap/classfile"

type FieldRef struct {
	MemberRef
	field *Field //缓存解析后的字段指针
}

//newFieldRef 创建FieldRef实例
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	/*如果类D想通过字段符号引用访问类C的某个字段，首先要解 析符号引用得到类C，然后根据字段名和描述符查找字段*/
	field := lookupField(c, self.name, self.descriptors)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

func lookupField(c *Class, name string, descriptors string) *Field {
	//首先在类C的字段中查找
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptors {
			return field
		}
	}
	//如果找不到，在类C的直接接口递归递归查找
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptors); field != nil {
			return field
		}
	}
	//如果还找不到的话，在类C的父类中递归查找
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptors)
	}
	return nil
}

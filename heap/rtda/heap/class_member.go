package heap

import "JVM-Go/heap/classfile"

//ClassMember 字段信息和方法信息都属于类的成员信息，抽取公共结构体
type ClassMember struct {
	accessFlags uint16 //访问标志
	name        string //类成名的名字(字段名或者方法名)
	descriptor  string //描述符(参数列表和返回值)
	class       *Class //字段或方法所属的类
}

//copyMemberInfo 从classFile文件中复制数据给Class实例
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.GetAccessFlags()
	self.name = memberInfo.GetName()
	self.descriptor = memberInfo.GetDescriptor()
}

//isAccessibleTo 字段或者方法的访问控制
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	/*如果字段是public，则任何类都可以访问*/
	if self.IsPublic() {
		return true
	}
	c := self.class
	/*如果字段是protected，则只有子类和同一个包下的类可以访问*/
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	/*如果字段有默认访问权限（非public，非protected，也 非privated），则只有同一个包下的类可以访问*/
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	/*Private 私有的，只有声明这个字段的类才能访问*/
	return d == c
}

func (self *ClassMember) IsPublic() bool {
	return self.accessFlags&ACC_PUBLIC != 0
}

func (self *ClassMember) IsPrivate() bool {
	return self.accessFlags&ACC_PRIVATE != 0
}

func (self *ClassMember) IsProtected() bool {
	return self.accessFlags&ACC_PROTECTED != 0
}

func (self *ClassMember) IsStatic() bool {
	return self.accessFlags&ACC_STATIC != 0
}

func (self *ClassMember) IsFinal() bool {
	return self.accessFlags&ACC_FINAL != 0
}

func (self *ClassMember) IsSynthetic() bool {
	return self.accessFlags&ACC_SYNTHETIC != 0
}

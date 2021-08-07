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

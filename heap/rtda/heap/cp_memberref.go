package heap

import "JVM-Go/heap/classfile"

//MemberRef 来存放字段和方法共有的信息
type MemberRef struct {
	SymRef
	name        string
	descriptors string
}

//copyMemberRefInfo 从classFile文件内存储的字段或方法常量中提取数据复制给MemberRef实例
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptors = refInfo.NameAndDescriptor()
}

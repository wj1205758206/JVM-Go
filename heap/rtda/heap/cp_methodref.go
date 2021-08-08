package heap

import "JVM-Go/heap/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func (self *MethodRef) GetName() string {
	return self.name
}

func (self *MethodRef) GetDescriptor() string {
	return self.descriptors
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

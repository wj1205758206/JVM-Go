package heap

import "JVM-Go/heap/classfile"

//ClassRef 类符号引用结构体，继承SymRef
type ClassRef struct {
	SymRef
}

//newClassRef 根据classFile文件中存储的类常量创建ClassRef实例
func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}

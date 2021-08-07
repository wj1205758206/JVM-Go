package heap

import "JVM-Go/heap/classfile"

//Field 字段信息的结构体，所有信息都是从ClassMember中继承过来的
type Field struct {
	ClassMember
}

//newFields 根据classFile文件的字段信息创建字段表
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fileds := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fileds[i] = &Field{}
		fileds[i].class = class
		fileds[i].copyMemberInfo(cfField)
	}
	return fileds
}

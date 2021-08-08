package heap

import "JVM-Go/heap/classfile"

//Field 字段信息的结构体，所有信息都是从ClassMember中继承过来的
type Field struct {
	ClassMember
	slotId          uint
	constValueIndex uint
}

//newFields 根据classFile文件的字段信息创建字段表
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fileds := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fileds[i] = &Field{}
		fileds[i].class = class
		fileds[i].copyMemberInfo(cfField)
		fileds[i].copyAttributes(cfField)
	}
	return fileds
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) IsVolatile() bool {
	return self.accessFlags&ACC_VOLATILE != 0
}

func (self *Field) IsTransient() bool {
	return self.accessFlags&ACC_TRANSIENT != 0
}

func (self *Field) IsEnum() bool {
	return self.accessFlags&ACC_ENUM != 0
}

func (self *Field) GetSlotId() uint {
	return self.slotId
}

func (self *Field) GetDescriptor() string {
	return self.descriptor
}

//copyAttributes 从字段属性表中读取constValueIndex
func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	constValueAttr := cfField.GetConstantValueAttribute()
	if constValueAttr != nil {
		self.constValueIndex = uint(constValueAttr.ConstantValueIndex())
	}
}

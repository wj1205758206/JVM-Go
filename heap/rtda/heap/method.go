package heap

import "JVM-Go/heap/classfile"

//Method 方法信息结构体，除了基继承ClassMember，还有一些额外信息
type Method struct {
	ClassMember
	maxStack  uint   //方法的操作数栈大小
	maxLocals uint   //方法的局部变量表大小
	code      []byte //关于方法的字节码
}

//newMethods 根据classFile文件中的方法信息创建Method表
func (self *Method) newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

//copyAttributes 从method_info结构中复制属性信息
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	codeAttr := cfMethod.GetCodeAttribute()
	if codeAttr != nil {
		self.maxStack = codeAttr.GetMaxStack()
		self.maxLocals = codeAttr.GetLocals()
		self.code = codeAttr.GetCode()
	}
}

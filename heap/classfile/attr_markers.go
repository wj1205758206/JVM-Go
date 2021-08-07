package classfile

//DeprecatedAttribute 起标记作用,用于指出类、接口、字段或方法已经不建议使用
type DeprecatedAttribute struct {
	MarkerAttribute
}

//SyntheticAttribute 起标记作用,用来标记源文件中不存在、由编译器生成的类成员
type SyntheticAttribute struct {
	MarkerAttribute
}
type MarkerAttribute struct {
	//attributeNameIndex
	//attributeLength 由于不包含任何数据，所以attribute_length的值必须是0
}

//readInfo 不包含数据所以该方法为空
func (self *MarkerAttribute) readInfo(reader *ClassReader) {

}

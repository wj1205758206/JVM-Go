package heap

/*SymRef 抽取类符号引用、字段符号引用、方法符号引用、接口方法符号引用的公共特性，封装成符号引用结构体
对于类符号引用，只要有类名，就可以解析符号引用
对于字段符号引用，首先要解析类符号引用得到类数据，然后用字段名和描述符查找字段数据
方法符号引用的解析过程和字段符号引用类似
*/
type SymRef struct {
	cp        *ConstantPool //存放符号引用所在的运行时常量池指针,可以通过符号引用访问到运行时常量池，进一步又可以访问到类数据
	className string        //存放类的完全限定名
	class     *Class        //缓存解析后的类结构体指针，这样类符号引用只需要解析一次就可以了，后续可以直接使用缓存值
}

//ResolvedClass 类符号引用解析
func (self *SymRef) ResolvedClass() *Class {
	//如果类符号引用已经解析，ResolvedClass方法直接返回类指针，否则调用resolveClassRef方法进行解析。
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

//resolveClassRef 进行类符号引用解析
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}

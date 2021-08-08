package heap

import (
	"JVM-Go/heap/classfile"
	"strings"
)

//Class 将要放进方法区内的类模板的结构体
type Class struct {
	accessFlags       uint16
	className         string        //类名
	superClassName    string        //父类名
	interfaceNames    []string      //接口名集合
	constantPool      *ConstantPool //运行时常量池指针
	fields            []*Field      //字段表
	methods           []*Method     //方法表
	loader            *ClassLoader  //存放类加载器指针
	superClass        *Class        //保存父类的指针
	interfaces        []*Class      //保存实现接口的指针集合
	instanceSlotCount uint          //存放实例变量占据空间的大小
	staticSlotCount   uint          //存放类变量占据空间的大小
	staticVars        Slots
}

//newClass 用来把ClassFile结构体转换成Class结构体
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.GetAccessFlags()
	class.className = cf.GetClassName()
	class.superClassName = cf.GetSuperClassName()
	class.interfaceNames = cf.GetInterfaceNames()
	class.constantPool = newConstantPool(class, cf.GetConstantPool())
	class.fields = newFields(class, cf.GetFields())
	class.methods = newMethods(class, cf.GetMethods())
	return class
}

//判断某个访问标志是否被设置

func (self *Class) IsPublic() bool {
	return self.accessFlags&ACC_PUBLIC != 0
}

func (self *Class) IsFinal() bool {
	return self.accessFlags&ACC_FINAL != 0
}

func (self *Class) IsSuper() bool {
	return self.accessFlags&ACC_SUPER != 0
}

func (self *Class) IsInterface() bool {
	return self.accessFlags&ACC_INTERFACE != 0
}

func (self *Class) IsAbstract() bool {
	return self.accessFlags&ACC_ABSTRACT != 0
}

func (self *Class) IsSynthetic() bool {
	return self.accessFlags&ACC_SYNTHETIC != 0
}

func (self *Class) IsAnnotation() bool {
	return self.accessFlags&ACC_ANNOTATION != 0
}

func (self *Class) IsEnum() bool {
	return self.accessFlags&ACC_ENUM != 0
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.className, "/"); i >= 0 {
		return self.className[:i]
	}
	return ""
}

func (self *Class) GetConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) GetStaticVars() Slots {
	return self.staticVars
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name string, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.GetName() == name && method.GetDescriptor() == descriptor {
			return method
		}
	}
	return nil
}

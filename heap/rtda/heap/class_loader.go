package heap

import (
	"JVM-Go/heap/classfile"
	"JVM-Go/heap/classpath"
	"fmt"
)

type ClassLoader struct {
	cp       *classpath.Classpath //依赖Classpath来搜索和读取class文件
	classMap map[string]*Class    //记录已经加载的类数据,类似于方法区的具体实现
}

//NewClassLoader 创建ClassLoader实例
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(className string) *Class {
	class, ok := self.classMap[className]
	if ok {
		return class
	}
	return self.loadNonArrayClass(className) //数组类和普通类有很大的不同，它的数据并不是来自class文件，而是由Java虚拟机在运行期间生成
}

//loadNonArrayClass 类的加载：读取字节码文件数据-->解析class文件数据-->链接
func (self *ClassLoader) loadNonArrayClass(className string) *Class {
	data, entry := self.readClass(className) //找到class文件并把数据读取到内存
	class := self.defineClass(data)          //解析class文件，生成虚拟机可以使用的类数据,并放入方法区
	link(class)                              //进行类的链接
	fmt.Printf("[Loaded %s from %v]\n", className, entry)
	return class
}

//link 类的链接
func link(class *Class) {
	verify(class)  //验证环节
	prepare(class) //准备环节
}

func verify(class *Class) {
	//TODO
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class) //计算实例字段的个数
	calcStaticFieldSlotIds(class)   //计算静态字段的个数
	allocAndInitStaticVars(class)   //给类变量分配空间，然后给它们赋予默认初始值
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			//如果静态变量属于基本类型或String类型，有final修饰符， 且它的值在编译期已知，则该值存储在class文件常量池中
			initStaticFinalVar(class, field)
		}
	}
}

//initStaticFinalVar 从常量池中加载常量值，然后给静态变量赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.GetConstantValueIndex()
	slotId := field.GetSlotId()
	if cpIndex > 0 {
		switch field.GetDescriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(slotId).(int32)
			vars.SetInt(cpIndex, val)
		case "J":
			val := cp.GetConstant(slotId).(int64)
			vars.SetLong(cpIndex, val)
		case "F":
			val := cp.GetConstant(slotId).(float32)
			vars.SetFloat(cpIndex, val)
		case "D":
			val := cp.GetConstant(slotId).(float64)
			vars.SetDouble(cpIndex, val)
		case "Ljava/lang/String;":
			panic("TODO")
		}
	}
}

//calcStaticFieldSlotIds 计算静态字段的个数
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

//calcInstanceFieldSlotIds 计算实例字段的个数
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

//readClass 调用了Classpath的ReadClass()方法，读取类路径下类的信息数据
func (self *ClassLoader) readClass(className string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(className)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + className)
	}
	return data, entry //返回读取的类加载信息数据和最终加载class文件的类路径
}

//defineClass 解析class文件数据
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data) //把class文件数据转换成Class结构体
	class.loader = self
	resolveSuperClass(class) //解析父类符号引用
	resolveInterfaces(class) //解析接口符号引用
	self.classMap[class.className] = class
	return class
}

//resolveInterfaces 解析接口符号引用
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName) //递归调用LoadClass()方法加载类的每一个直接接口
		}
	}
}

//resolveSuperClass 解析父类符号引用
func resolveSuperClass(class *Class) {
	if "java/lang/Object" != class.className {
		class.superClass = class.loader.LoadClass(class.superClassName) //有父类需要加载父类
	}
}

//parseClass 把class文件数据转换成Class结构体
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data) //将classFile文件数据转换成classFile结构体
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf) //将classFile结构体转换成Class结构体保存在方法区
}

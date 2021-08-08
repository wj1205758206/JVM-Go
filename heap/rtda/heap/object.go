package heap

type Object struct {
	class  *Class //存放对象的Class指针
	fields Slots  //存放实例变量
}

func (self *Object) GetFields() Slots {
	return self.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

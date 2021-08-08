package heap

//isAssignableFrom
func (self *Class) isAssignableFrom(other *Class) bool {
	//在三种情况下，S类型的引用值可以赋值给T类型
	s, t := other, self
	if s == t { //S 和T是同一类型
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t) //T是类且S是T的子类
	} else {
		return s.isImplements(t) //T是接口且S实现了T接口
	}
}

//isSubClassOf 判断是否是子类
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

//isImplements 判断是否实现了其接口
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

//isSubInterfaceOf 判断是否实现了其子接口
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

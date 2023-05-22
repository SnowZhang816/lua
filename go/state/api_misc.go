package state

// import "fmt"
// import "main/api"

func (self *luaState) Len(idx int) {
	val := self.stack.get(idx)
	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else if result, ok := callMetaMethod(val, val, "__len", self); ok { 
		self.stack.push(result)
	} else if t, ok := val.(*luaTable); ok {
		self.stack.push(int64(t.len()))
	} else {
		panic("length error!")
	}
}

func (self *luaState) RawLen(idx int) {
	val := self.stack.get(idx)
	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else if t, ok := val.(*luaTable); ok {
		self.stack.push(int64(t.len()))
	} else {
		panic("length error!")
	}
}

func (self *luaState) Concat(n int) {
	if n == 0 {
		self.stack.push("")
	} else if n >= 2 {
		for i := 1; i < n; i++{
			if self.IsString(-1) && self.IsString(-2) {
				s2 := self.ToString(-1)
				s1 := self.ToString(-2)
				self.stack.pop()
				self.stack.pop()
				self.stack.push(s1 + s2)
				continue
			}

			b := self.stack.pop()
			a := self.stack.pop()
			if result,ok := callMetaMethod(a, b, "__concat", self); ok {
				self.stack.push(result)
				continue
			}

			panic("concatenation error!")
		}
	}
}

func (self *luaState) Next(idx int) bool {
	val := self.stack.get(idx)
	if t, ok := val.(*luaTable); ok {
		key := self.stack.pop()
		if nextKey := t.nextKey(key); nextKey != nil {
			self.stack.push(nextKey)
			self.stack.push(t.get(nextKey))
			return true
		}
		return false
	}

	panic("table expected!")
}

func (self *luaState) Error() int {
	err := self.stack.pop()
	panic(err)
}
package state

import "main/api"

func (self *luaState) setTable(t,k,v luaValue, raw bool) {
	if tbl,ok := t.(*luaTable); ok {
		if raw || tbl.get(k) != nil || !tbl.hasMetaField("__newindex") {
			tbl.put(k,v)
			return
		}
	}

	if !raw {
		if mf := getMetaField(t, "__newindex", self); mf != nil {
			switch x := mf.(type) {
			case *luaTable:
				self.setTable(x, k, v, false)
				return
			case *closure:
				self.stack.push(mf)
				self.stack.push(t)
				self.stack.push(k)
				self.stack.push(v)
				self.Call(3, 0)
				return
			}
		}
	}

	panic("setTable index errotr!")
}

func (self *luaState) SetTable(idx int) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	k := self.stack.pop()
	self.setTable(t, k, v, false)
}

func (self *luaState) RawSet(idx int) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	k := self.stack.pop()
	self.setTable(t, k, v, true)
}

func (self *luaState) SetField(idx int, k string) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	self.setTable(t, k, v, false)
}

func (self *luaState) SetI(idx int, i int64) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	self.setTable(t, i, v, false)
}

func (self *luaState) RawSetI(idx int, i int64) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	self.setTable(t, i, v, true)
}

func (self *luaState) SetGlobal(name string) {
	t := self.registry.get(api.LUA_RIDX_GLOBALS)
	v := self.stack.pop()
	self.setTable(t, name, v, true)
}

func (self *luaState) Register(name string, f api.GoFunction) {
	self.PushGoFunction(f, 0)
	self.SetGlobal(name)
}

func (self *luaState) SetMetaTable(idx int) {
	val := self.stack.get(idx)
	tbl := self.stack.pop()
	if tbl == nil {
		setMateTable(val, nil, self)
	} else if mt, ok := tbl.(*luaTable); ok {
		setMateTable(val, mt, self)
	} else {
		panic("SetMetaTable table expected!")
	}
}

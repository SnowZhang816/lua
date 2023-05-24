package state

import "main/api"
import "main/cLog"

func (self *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	self.stack.push(t)
}

func (self *luaState) NewTable() {
	t := newLuaTable(0, 0)
	self.stack.push(t)
}

func (self *luaState) getTable(t,k luaValue, raw bool) api.LuaType {
	// cLog.Println("getTable", t, k)
	if tbl,ok := t.(*luaTable); ok {
		// cLog.Println("getTable", k)
		// tbl.printTable()
		v := tbl.get(k)
		if raw || v != nil || !tbl.hasMetaField("__index") {
			// cLog.Println("getTable", v)
			self.stack.push(v)
			return typeOf(v)
		}
	}

	if !raw {
		if mf := getMetaField(t, "__index", self); mf != nil {
			switch x := mf.(type) {
			case *luaTable:
				return self.getTable(x, k, false)
			case *closure:
				self.stack.push(mf)
				self.stack.push(t)
				self.stack.push(k)
				self.Call(2, 1)
				v := self.stack.get(-1)
				return typeOf(v)
			}
		}
	}

	panic("getTable index error!")
}

func (self *luaState) GetTable(idx int) api.LuaType {
	t := self.stack.get(idx)
	k := self.stack.pop()
	return self.getTable(t, k, false)
}

func (self *luaState) RawGet(idx int) api.LuaType {
	t := self.stack.get(idx)
	k := self.stack.pop()
	return self.getTable(t, k, true)
}


func (self *luaState) GetField(idx int, k string) api.LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, k, false)

	// self.PushString(k)
	// return self.GetTable(idx)
}

func (self *luaState) GetI(idx int, i int64) api.LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, i, false)
}

func (self *luaState) RawGetI(idx int, i int64) api.LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, i, true)
}

func (self *luaState) GetGlobal(name string) api.LuaType {
	t := self.registry.get(api.LUA_RIDX_GLOBALS)
	return self.getTable(t, name, true)
}

func (self *luaState) GetMetaTable(idx int) bool {
	val := self.stack.get(idx)

	if mt := getMateTable(val, self); mt != nil {
		self.stack.push(mt)
		return true
	}

	return false
}

func (self *luaState) PrintTable(idx int) {
	t := self.stack.get(idx)
	if tbl,ok := t.(*luaTable); ok {
		tbl.printTable()
	} else {
		cLog.Println("PrintTable not a table")
	}
}

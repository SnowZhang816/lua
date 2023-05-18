package state

import "main/api"
import "fmt"

func (self *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	self.stack.push(t)
}

func (self *luaState) NewTable() {
	t := newLuaTable(0, 0)
	self.stack.push(t)
}

func (self *luaState) getTable(t,k luaValue) api.LuaType {
	fmt.Println("getTable", t, k)
	if tbl,ok := t.(*luaTable); ok {
		v := tbl.get(k)
		fmt.Println("getTable", v)
		self.stack.push(v)
		
		return typeOf(v)
	}

	panic("not a table!")
}

func (self *luaState) GetTable(idx int) api.LuaType {
	t := self.stack.get(idx)
	k := self.stack.pop()
	return self.getTable(t, k)
}

func (self *luaState) GetField(idx int, k string) api.LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, k)

	// self.PushString(k)
	// return self.GetTable(idx)
}

func (self *luaState) GetI(idx int, i int64) api.LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, i)
}

func (self *luaState) PrintTable(idx int) {
	t := self.stack.get(idx)
	if tbl,ok := t.(*luaTable); ok {
		fmt.Println("PrintTable", tbl)
	} else {
		fmt.Println("PrintTable not a table")
	}
}

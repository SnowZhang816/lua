package state

import "math"
import "main/number"
import "main/cLog"

type luaTable struct {
	metaTable *luaTable
	arr []luaValue
	_map map[luaValue]luaValue
	keys map[luaValue]luaValue
}

func newLuaTable(nArr, nRec int) *luaTable {
	t := &luaTable{}
	if nArr > 0 {
		t.arr = make([]luaValue, 0, nArr)
	}

	if nRec > 0 {
		t._map = make(map[luaValue]luaValue, nRec)
	}
	return t
}

func _floatToInteger(key luaValue) luaValue {
	if k, ok := key.(float64); ok {
		if i, ok := number.FloatToInteger(k); ok {
			return i
		}
	}
	return key
}

func (self *luaTable) get(key luaValue) luaValue {
	k := _floatToInteger(key)
	if idx, ok := k.(int64); ok {
		if idx >= 1 && idx <= int64(len(self.arr)) {
			return self.arr[idx - 1]
		}
	}
	return self._map[k]
}

func (self *luaTable) _shrinkArray()  {
	for i := len(self.arr) - 1; i >=0; i-- {
		if self.arr[i] == nil {
			self.arr = self.arr[0:i]
		}
	}
}

func (self *luaTable)  _expandArray()  {
	for idx := int64(len(self.arr)) + 1; true; idx++ {
		if val, found := self._map[idx]; found {
			delete(self._map, idx)
			self.arr = append(self.arr, val)
		} else {
			break
		}
	}
}

func (self *luaTable) put(key,val luaValue) {
	if key == nil {
		panic("table index is nil!")
	}
	if f,ok := key.(float64); ok && math.IsNaN(f) {
		panic("table index is NaN!")
	}

	k := _floatToInteger(key)

	if idx,ok := k.(int64); ok && idx > 0 {
		arrLen := int64(len(self.arr))
		if idx <= arrLen {
			self.arr[idx - 1] = val
			if idx == arrLen && val == nil {
				self._shrinkArray()
			}
			return
		}

		if idx == arrLen + 1 {
			delete(self._map, key)
			if val != nil {
				self.arr = append(self.arr, val)
				self._expandArray()
			}
			return
		}
	}

	if val != nil {
		if self._map == nil {
			self._map = make(map[luaValue]luaValue, 8)
		}
		self._map[key] = val
	} else {
		delete(self._map, key)
	}
}

func (self *luaTable) len() int {
	return len(self.arr)
}

func (self *luaTable) hasMetaField(fieldName string) bool {
	return self.metaTable != nil && self.metaTable.get(fieldName) != nil
}

func (self *luaTable) initKeys() {
	self.keys = make(map[luaValue]luaValue, )

	var key luaValue = nil
	for i, v := range self.arr {
		if v != nil {
			self.keys[key] = int64(i + 1)
			key = int64(i + 1)
		}
	}

	for k,v := range self._map {
		if v != nil {
			self.keys[key] = k
			key = k
		}
	}

	cLog.Println("initKeys", self.keys)
}

func (self *luaTable) nextKey(key luaValue) luaValue {
	if self.keys == nil || key == nil {
		self.initKeys()
		// self.changed = false
	}

	return self.keys[key]
}

func (self *luaTable) printTable() {
	// cLog.Println(self)
	cLog.Print("table@")
	cLog.Printf("%p", &self)
	cLog.Print("[")
	cLog.Printf("metaTable:")
	cLog.Print(self.metaTable)
	cLog.Printf(" arr(%d):", len(self.arr))
	cLog.Print(self.arr)
	cLog.Printf(" _map:")
	cLog.Print(self._map)
	cLog.Print("]")
	cLog.Println()
}
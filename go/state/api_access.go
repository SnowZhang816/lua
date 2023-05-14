package state

import "fmt"
import "main/api"

const (
	LUA_TNONE = iota - 1
	LUA_TNIL
	LUA_TBOOLEAN
	LUA_TLIGHTUSRDATA
	LUA_TNUMBER
	LUA_TSTRING
	LUA_TTABLE
	LUA_TFUNCTION
	LUA_TUSERDATA
	LUA_TTHREAD
)

func (self *luaState) TypeName(tp api.LuaType) string {
	switch tp {
	case LUA_TNONE:					return "no value"
	case LUA_TNIL:					return "nil"
	case LUA_TBOOLEAN:				return "boolean"
	case LUA_TLIGHTUSRDATA:			return "light userdata"
	case LUA_TNUMBER:				return "number"
	case LUA_TSTRING:				return "string"
	case LUA_TTABLE:				return "table"
	case LUA_TFUNCTION:				return "function"
	case LUA_TUSERDATA:				return "userdata"
	case LUA_TTHREAD:				return "thread"
	default:						return ""
	}
}

func (self *luaState) Type(idx int) api.LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

func (self *luaState) IsNone(idx int) bool {
	return self.Type(idx) == LUA_TNONE
}

func (self *luaState) IsNil(idx int) bool {
	return self.Type(idx) == LUA_TNIL
}

func (self *luaState) IsNoneOrNil(idx int) bool {
	return self.Type(idx) <= LUA_TNIL
}

func (self *luaState) IsBoolean(idx int) bool {
	return self.Type(idx) == LUA_TBOOLEAN
}

func (self *luaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

func (self *luaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

func (self *luaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (self *luaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

func (self *luaState) ToNumber(idx int) float64 {
	n, _ := self.ToNumberX(idx)
	return n
}

func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case float64:	return x, true
	case int64:		return float64(x), true
	default:		return 0, false
	}
}

func (self *luaState) ToInteger(idx int) int64 {
	i, _ := self.ToIntegerX(idx)
	return i
}

func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

func (self *luaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}

func (self *luaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case string:				return x, true
	case int64, float64:		
		s := fmt.Sprintf("%v", x)
		self.stack.set(idx, s)
		return s, true
	default:					return "", false
	}
}
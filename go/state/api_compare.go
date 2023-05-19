package state

import "main/api"

func (self *luaState) Compare(idx1, idx2 int, op api.CompareOP) bool {
	a := self.stack.get(idx1)
	b := self.stack.get(idx2)
	switch op {
	case api.LUA_OPEQ:		return _eq(a, b, self)
	case api.LUA_OPLT:		return _lt(a, b, self)
	case api.LUA_OPLE:		return _le(a, b, self)
	default: panic("invalid compare op!")
	}
}

func (self *luaState) RawEqual(idx1,idx2 int) bool {
	a := self.stack.get(idx1)
	b := self.stack.get(idx2)
	return _eq(a, b, nil)
}

func _eq(a, b luaValue, ls *luaState) bool {
	switch x := a.(type) {
	case nil:
		return b == nil
	case bool:
		y, ok := b.(bool)
		return ok && x == y
	case int64:
		switch y := b.(type) {
		case int64: 	return x == y
		case float64:	return float64(x) == y
		default:		return false
		}
	case float64:
		switch y := b.(type) {
			case int64: 	return x == float64(y)
			case float64:	return x == y
			default:		return false
			}
	case string:
		y, ok := b.(string)
		return ok && x == y
	case *luaTable:
		if y,ok := b.(*luaTable); ok && x != y && ls != nil {
			if result, ok := callMetaMethod(x,y,"__eq",ls); ok {
				return convertToBoolean(result)
			}
		}
		return a == b
	default:
		return a == b
	}
}

func _lt(a, b luaValue, ls *luaState) bool {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64: 	return x < y
		case float64:	return float64(x) < y
		}
	case float64:
		switch y := b.(type) {
		case int64: 	return x < float64(y)
		case float64:	return x < y
		}
	case string:
		y, ok := b.(string)
		if ok {
			return x < y
		}
	}

	if result,ok := callMetaMethod(a, b, "__lt", ls); ok {
		return convertToBoolean(result)
	} else {
		panic("comparison error!")
	}
}

func _le(a, b luaValue, ls *luaState) bool {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64: 	return x <= y
		case float64:	return float64(x) <= y
		}
	case float64:
		switch y := b.(type) {
		case int64: 	return x <= float64(y)
		case float64:	return x <= y
		}
	case string:
		y, ok := b.(string)
		if ok {
			return x <= y
		}
	}
	if result,ok := callMetaMethod(a, b, "__le", ls); ok {
		return convertToBoolean(result)
	} else if result,ok := callMetaMethod(b, a, "__lt", ls); ok {
		return convertToBoolean(result)
	} else {
		panic("comparison error!")
	}
}

package stdlib

import "main/api"
import "main/cLog"
import "sort"
import "strings"

type wrapper struct {
	ls api.LuaState
}

func (self wrapper)Len() int {
	return int(self.ls.Len2(1))
}

func (self wrapper)Less(i,j int) bool {
	ls := self.ls
	if ls.IsFunction(2) {
		ls.PushValue(2)
		ls.GetI(1, int64(i + 1))
		ls.GetI(1, int64(j + 1))
		ls.Call(2, 1)
		b := ls.ToBoolean(-1)
		ls.Pop(1)
		return b
	} else {
		ls.GetI(1, int64(i + 1))
		ls.GetI(1, int64(j + 1))
		b := ls.Compare(-2, -1, api.LUA_OPLT)
		ls.Pop(2)
		 return b
	}
}

func (self wrapper)Swap(i,j int) {
	ls := self.ls
	ls.GetI(1, int64(i + 1))
	ls.GetI(1, int64(j + 1))
	ls.SetI(1, int64(i + 1))
	ls.SetI(1, int64(j + 1))
}

var tableLib = map[string]api.GoFunction{
	// "move":			move,
	// "insert":		insert,
	// "remove":		remove,
	"sort":			tabSort,
	"concat":		tabConcat,
	// "pack":			pack,
	// "unpack":		unpack,
}

// func move(ls api.LuaState) int {
// }

// func insert(ls api.LuaState) int {
// }

// func remove(ls api.LuaState) int {
// }

func tabSort(ls api.LuaState) int {
	sort.Sort(wrapper{ls})
	return 0
}

func tabConcat(ls api.LuaState) int {
	tabLen := ls.Len2(1)
	step := ls.OptString(2, "")
	i := ls.OptInteger(3, 1)
	j := ls.OptInteger(3, tabLen)

	if i > j {
		ls.PushString("")
	}

	buf := make([]string, j - i + 1)

	for k := i; k <= j; k++ {
		ls.GetI(1, k)
		if !ls.IsString(-1) {
			ls.Error2("invalid value (%s) at index (%d) in table for concat", ls.TypeName2(-1), k)
		}

		buf[k - 1] = ls.ToString(-1)
		ls.Pop(1)
	}

	ls.PushString(strings.Join(buf, step))
	return 1
}

// func pack(ls api.LuaState) int {
// }

// func unpack(ls api.LuaState) int {
// }

func OpenTableLib(ls api.LuaState) int {
	cLog.Println("OpenTableLib")
	ls.NewLib(tableLib)
	return 1
}
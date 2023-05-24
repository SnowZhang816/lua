package stdlib

// import "os"
import "time"
import "main/cLog"
import "main/api"

var osLib = map[string]api.GoFunction{
	"time":			osTime,
	// "date":			osDate,
}

func _getField(ls api.LuaState, key string, dft int64) int {
	t := ls.GetField(-1, key)
	res, isNum := ls.ToIntegerX(-1)
	if !isNum {
		if t != api.LUA_TNIL {
			return ls.Error2("field '%s' is not a integer", key)
		} else if dft < 0 {
			return ls.Error2("field '%s' missing in date table", key)
		}
		res = dft
	}

	ls.Pop(1)
	return int(res)
}

func osTime(ls api.LuaState) int {
	if ls.IsNoneOrNil(1) {
		t := time.Now().Unix()
		ls.PushInteger(t)
	} else {
		ls.CheckType(1, api.LUA_TTABLE)
		sec := _getField(ls, "sec", 0)
		min := _getField(ls, "min", 0)
		hour := _getField(ls, "hour", 12)
		day := _getField(ls, "day", -1)
		month := _getField(ls, "month", -1)
		year := _getField(ls, "year", -1)

		t := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local).Unix()
		ls.PushInteger(t)
	}

	return 1
}

// func osDate(ls api.LuaState) int {

// }

func OpenOsLib(ls api.LuaState) int {
	cLog.Println("OpenOsLib")
	ls.NewLib(osLib)
	return 1
}
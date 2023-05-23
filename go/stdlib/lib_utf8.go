package stdlib

import "unicode/utf8"
import "main/cLog"
import "main/api"

var utf8Lib = map[string]api.GoFunction{
	// "len":			utfLen,
	// "offset":		utfOffset,
	// "codepoint":	utfCodePoint,
	"char":			utfChar,
	// "codes":		utfCodes,
}

func _encodeUtf8(codePoints []rune) string {
	buf := make([]byte, 6)
	str := make([]byte, 0, len(codePoints))

	for _, cp := range codePoints {
		n := utf8.EncodeRune(buf, cp)
		str = append(str, buf[0:n]...)
	} 

	return string(str)
}

func utfChar(ls api.LuaState) int {
	n := ls.GetTop()
	codePoints := make([]rune, n)

	for i := 1; i <= n; i ++ {
		cp := ls.CheckInteger(i)
		ls.ArgCheck(0 <= cp && cp <= 0x10FFFF, i, "value out of range")
		codePoints[i - 1] = rune(cp)
	}
	ls.PushString(_encodeUtf8(codePoints))
	return 1
}

func OpenUtf8Lib(ls api.LuaState) int {
	cLog.Println("OpenUtf8Lib")
	ls.NewLib(utf8Lib)
	return 1
}
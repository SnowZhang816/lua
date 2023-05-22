package vm

import "main/cLog"
import "math"

/*
**converts an integer to a "floating point byte", represented as
**(eeeeexxx), where the real value is (1xxx) * 2^(eeeee - 1) if 
**(eeeee != 0) and (xxx) otherwise
*/
func Int2fb(x int) int {
	e := 0	/*exponent*/
	if x < 8 {
		return x
	}
	cLog.Printf("%b-%d\n", x, x)
	for x >= (8 << 4) { /* coarse steps */
		x = (x + 0xf) >> 4 /* x = ceil(x / 16) */
		cLog.Printf("coarse-%b-%d\n", x, x)
		e += 4
	}
	for x >= (8 << 1) { /* fine steps */
		x = (x + 1) >> 1 /* x ceil(x / 2) */
		cLog.Printf("fine-%b-%d\n", x, x)
		e++
	}
	cLog.Printf("finally-%d\n", e)
	return ((e + 1) << 3) | (x - 8)
}

func Int2fb1(x int) int {
	e := 0	/*exponent*/
	if x < 8 {
		return x
	}
	cLog.Printf("%b-%d\n", x, x)
	for x >= 128 {
		x = int(math.Ceil(float64(x)/16))
		cLog.Printf("coarse+%b-%d\n", x, x)
		e += 4
	}

	for x >= 16 { /* fine steps */
		x = int(math.Ceil(float64(x) / 2))
		cLog.Printf("coarse+%b-%d\n", x, x)
		e++
	}
	cLog.Printf("finally-%d\n", e)
	return (e+1)*int(math.Pow(2,3)) + (x - 8)
}
// 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,18,20,22,24,26,28,30,32,36,40
// 10000,10001,10010,10011,10100,...,11111,100000,100001,100010
/*converts back*/
func Fb2int(x int) int {
	if x < 8 {
		return x
	} else {
		return ((x & 7) + 8) << uint((x >> 3) - 1)
	}
}
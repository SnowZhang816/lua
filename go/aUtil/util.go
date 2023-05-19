package aUtil

import "fmt"
import "main/binChunk"
import "main/vm"

func PrintProto(f *binChunk.Prototype) {
	printHeader(f)
	printCode(f)
	printDetail(f)
 
	for _,p := range f.Protos {
		PrintProto(p)
	}
 }

func printHeader(f *binChunk.Prototype) {
	funcType := "main"
	if f.LineDefined > 0 {
	   funcType = "function"
	}
 
	varArgFlag := ""
	if f.IsVarArg > 0 {
	   varArgFlag = "+"
	}
 
	fmt.Printf("\n%s <%s:%d,%d> (%d instruction)\n", funcType, f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))
	fmt.Printf("%d%s params, %d slots, %d upvalues, ", f.NumParams, varArgFlag, f.MaxStackSize, len(f.UpValues))
	fmt.Printf("%d locals, %d constants, %d functions\n", len(f.LocVars), len(f.Constants), len(f.Protos))
}

func printCode(f *binChunk.Prototype) {
	// fmt.Println("printCode",f.Code,f.LineInfo)
	for pc, c := range f.Code {
	   line := "-"
	   if len(f.LineInfo) > 0 {
		  line = fmt.Sprintf("%d", f.LineInfo[pc])
	   }
 
	   i := vm.Instruction(c)
	   fmt.Printf("\t%d\t[%s]\t%s \t", pc + 1, line, i.OpName())
	   printOperands(i)
	   fmt.Println("\n")
	}
}

func printOperands(i vm.Instruction) {
	switch i.OpMode() {
	case vm.IABC:
	   a, b, c := i.ABC()
	   fmt.Printf("%d", a)
	   if i.BMode() != vm.OpArgN {
		  if b > 0xFF {
			 fmt.Printf(" %d", -1-b&0xFF)
		  } else {
			 fmt.Printf(" %d", b)
		  }
	   }
	   
	   if i.CMode() != vm.OpArgN {
		  if c > 0xFF {
			 fmt.Printf(" %d", -1-c&0xFF)
		  } else {
			 fmt.Printf(" %d", c)
		  }
	   }
	case vm.IABx:
	   a, bx := i.ABx()
	   fmt.Printf("%d", a)
	   if i.BMode() == vm.OPArgK {
		  fmt.Printf(" %d", -1-bx)
	   } else if i.BMode() == vm.OpArgU {
		  fmt.Printf(" %d", bx)
	   }
	case vm.IAsBx:
	   a, sbx := i.AsBx()
	   fmt.Printf("%d %d", a, sbx)
	case vm.IAx:
	   ax := i.Ax()
	   fmt.Printf("%d", -1 - ax)
	}
}

func printDetail(f *binChunk.Prototype) {
	fmt.Printf("constants (%d):\n", len(f.Constants))
	for i, k := range f.Constants {
	   fmt.Printf("\t%d\t%s\n", i+1, constantsToString(k))
	}
 
	fmt.Printf("locals (%d):\n", len(f.LocVars))
	for i, locVar := range f.LocVars {
	   fmt.Printf("\t%d\t%s\t\t%d\t%d\n", i, locVar.VarName, locVar.StartPC, locVar.EndPC)
	}
 
	fmt.Printf("upValues (%d):\n", len(f.UpValues))
	for i, upValue := range f.UpValues {
	   fmt.Printf("\t%d\t%s\t%d\t%d\n", i, upValName(f, i), upValue.InStack, upValue.Idx)
	}
}

func constantsToString(k interface{}) string {
	switch k.(type){
	case nil:         return "nil"
	case bool:        return fmt.Sprintf("%t", k)
	case float64:     return fmt.Sprintf("%g", k)
	case int64:       return fmt.Sprintf("%d", k)
	case string:      return fmt.Sprintf("%q", k)
	default:          return "?"
	}
}
 
 func upValName(f *binChunk.Prototype, idx int) string {
	if len(f.UpValuesNames) > 0 {
	   return f.UpValuesNames[idx]
	}
	return "-"
}

func GoPrint(s string) {
	
}
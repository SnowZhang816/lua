package aUtil

import "main/cLog"
import "main/binChunk"
import "main/vm"
import "fmt"

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
 
	cLog.Printf("\n%s <%s:%d,%d> (%d instruction)\n", funcType, f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))
	cLog.Printf("%d%s params, %d slots, %d upvalues, ", f.NumParams, varArgFlag, f.MaxStackSize, len(f.UpValues))
	cLog.Printf("%d locals, %d constants, %d functions\n", len(f.LocVars), len(f.Constants), len(f.Protos))
}

func printCode(f *binChunk.Prototype) {
	// cLog.Println("printCode",f.Code,f.LineInfo)
	for pc, c := range f.Code {
	   line := "-"
	   if len(f.LineInfo) > 0 {
		  line = fmt.Sprintf("%d", f.LineInfo[pc])
	   }
 
	   i := vm.Instruction(c)
	   cLog.Printf("\t%d\t[%s]\t%s \t", pc + 1, line, i.OpName())
	   printOperands(i)
	   cLog.Println("\n")
	}
}

func printOperands(i vm.Instruction) {
	switch i.OpMode() {
	case vm.IABC:
	   a, b, c := i.ABC()
	   cLog.Printf("%d", a)
	   if i.BMode() != vm.OpArgN {
		  if b > 0xFF {
			 cLog.Printf(" %d", -1-b&0xFF)
		  } else {
			 cLog.Printf(" %d", b)
		  }
	   }
	   
	   if i.CMode() != vm.OpArgN {
		  if c > 0xFF {
			 cLog.Printf(" %d", -1-c&0xFF)
		  } else {
			 cLog.Printf(" %d", c)
		  }
	   }
	case vm.IABx:
	   a, bx := i.ABx()
	   cLog.Printf("%d", a)
	   if i.BMode() == vm.OPArgK {
		  cLog.Printf(" %d", -1-bx)
	   } else if i.BMode() == vm.OpArgU {
		  cLog.Printf(" %d", bx)
	   }
	case vm.IAsBx:
	   a, sbx := i.AsBx()
	   cLog.Printf("%d %d", a, sbx)
	case vm.IAx:
	   ax := i.Ax()
	   cLog.Printf("%d", -1 - ax)
	}
}

func printDetail(f *binChunk.Prototype) {
	cLog.Printf("constants (%d):\n", len(f.Constants))
	for i, k := range f.Constants {
	   cLog.Printf("\t%d\t%s\n", i+1, constantsToString(k))
	}
 
	cLog.Printf("locals (%d):\n", len(f.LocVars))
	for i, locVar := range f.LocVars {
	   cLog.Printf("\t%d\t%s\t\t%d\t%d\n", i, locVar.VarName, locVar.StartPC, locVar.EndPC)
	}
 
	cLog.Printf("upValues (%d):\n", len(f.UpValues))
	for i, upValue := range f.UpValues {
	   cLog.Printf("\t%d\t%s\t%d\t%d\n", i, upValName(f, i), upValue.InStack, upValue.Idx)
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
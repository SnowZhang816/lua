package main

import (
   "fmt"
   "main/state"
   "io/ioutil"
   "os"
   "unsafe"
   "bytes"
   "encoding/binary"
)

import "main/api"

func main() {
   fmt.Println("Hello World!")
   checkEndian()

   if len(os.Args) > 1 {
      data, err := ioutil.ReadFile(os.Args[1])
      if err != nil {
         panic(err)
      }

      ls := state.New()
      ls.Register("print", print)
      ls.Load(data, os.Args[1], "b")
      ls.Call(0, 0)
   }
}

//整形转换成字节
func IntToBytes(n int) []byte {
   x := int32(n)

   bytesBuffer := bytes.NewBuffer([]byte{})
   binary.Write(bytesBuffer, binary.BigEndian, x)
   return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
   bytesBuffer := bytes.NewBuffer(b)

   var x int32
   binary.Read(bytesBuffer, binary.BigEndian, &x)

   return int(x)
}

func checkEndian() {
   var value int32 = 1 // 占4byte 转换成16进制 0x00 00 00 01 
   // 大端(16进制)：00 00 00 01
   // 小端(16进制)：01 00 00 00
   pointer := unsafe.Pointer(&value)
   pb := (*byte)(pointer)
   if *pb != 1{
     fmt.Println("大端模式")
   }
   fmt.Println("小端模式")
}

func print(ls api.LuaState) int {
   nArgs := ls.GetTop()
   fmt.Printf("LUAPrint:")
   for i := 1; i <= nArgs; i++ {
      if ls.IsBoolean(i) {
         fmt.Printf("%t", ls.ToBoolean(i))
      } else if ls.IsString(i) {
         fmt.Printf("%s", ls.ToString(i))
      } else {
         fmt.Printf("%s",ls.TypeName(ls.Type(i)))
      }
      if i < nArgs {
         fmt.Print(" ")
      }
   }
   fmt.Println()
   return 0
}

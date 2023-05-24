package main

import (
   "main/state"
   "os"
   "unsafe"
   "bytes"
   "encoding/binary"
   "main/cLog"
)

func main() {
   cLog.Println("Hello World!","dsadsa","dasdas")
   checkEndian()

   if len(os.Args) > 1 {
      ls := state.New()
      ls.OpenLibs()

      ls.PrintStack(true)
      ls.PrintGlobalTable()
      ls.PrintRegister()
      ls.PrintLoadedTable()
      
      ls.LoadFile(os.Args[1])
      ls.Call(0, -1)
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
     cLog.Println("大端模式")
   }
   cLog.Println("小端模式")
}




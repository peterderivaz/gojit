// Build with
//  go build -o bin ./...
//  gdb bin/calc
//  run
//  disassemble $pc,+10
package main

import (
    "github.com/peterderivaz/gojit"
    . "github.com/peterderivaz/gojit/amd64"
    "fmt"
)

func main() {
  for x := 0; x < 10;x++ {
    asm, err := NewGoABI(gojit.PageSize)
    if err != nil {
        panic(err)
    }
    //asm.Mov(Imm{int32(x)},Eax)
    asm.Mov(Imm{int32(-0x80000000)},Eax)
    //asm.Mov(Indirect{Rbx, 0, 32},Eax)
    asm.Sar(Imm{int32(x)},Eax) // src,dst
    asm.Mov(Eax,Indirect{Rbx, 0, 32})
    //asm.SegFault()
    asm.Ret()
    var f1 func()
    asm.BuildTo(&f1)
    
    gojit.JitData[0] = 100*uint32(x)
    f1()
    fmt.Printf("%x->%x\n",x,gojit.JitData[0])
    
    asm.Release()
  }  
}

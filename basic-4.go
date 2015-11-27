package main

import (
    "fmt"
    "strconv"
)

const(
    a=1 << (10 * iota)
    b
    c
    d
)


func 打印(s int){
    fmt.Println(strconv.Itoa(s))
}

func main(){
    打印(a)
    打印(b)
    打印(c)
    打印(d)
}

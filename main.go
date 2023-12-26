package main

import (
	"fmt"
	"log"
)

func main(){
    str := "the quick brown fox jumped over the lazy white dog"
    pt := newPieceTable(str)
    if err := pt.add("went to the park and ", 20);err!=nil{
        log.Fatal(err)
    }
    if err := pt.add(".", 71);err!=nil{
        log.Fatal(err)
    }
    if err := pt.delete(41, 7);err!=nil{
        log.Fatal(err)
    }
    fmt.Println(pt.toString())
}

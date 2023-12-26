package main

import "fmt"

func main(){
    str := "the quick brown fox jumped over the lazy white dog"
    pt := newPieceTable(str)
    pt.add("went to the park and ", 19)
    pt.add(".", 44)
    pt.delete(41, 6)
    fmt.Println(pt.toString())
}

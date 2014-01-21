package main

import (
    "fmt"
    "time"
    "strconv"
)

func main (){
    t := time.Now ().Unix ()
    var i int
    i = 10
    fmt.Println (strconv.FormatInt (i, 10))
    fmt.Println (t)
}

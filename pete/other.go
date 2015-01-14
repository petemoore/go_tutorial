package main

import (
    "fmt"
//    "github.com/petemoore/go_tutorial/lib"
)

type fred struct {
    x string
    y int
}

type tom struct {
    f fred
    j uint8
}

func main() {
    t := tom{j: 3, f: fred{x: "yo", y: 5}}
    fmt.Println("Pointer/Value differences w.r.t. if function is defined against pointer or value")
    fmt.Println("================================================================================")
    fmt.Printf("Initial value: %v\n", t.j)
    t.no_play()
    fmt.Printf("When function is defined against a value: %v\n", t.j)
    t.play_with()
    fmt.Printf("When function is defined against a pointer: %v\n", t.j)

    fmt.Println("Pointer/Value differences when passed as a parameter to a function")
    fmt.Println("==================================================================")
    fmt.Printf("Initial value: %v\n", t.j)
    no_fiddle(t)
    fmt.Printf("When function is passed a value: %v\n", t.j)
    fiddle(&t)
    fmt.Printf("When function is passed a pointer: %v\n", t.j)
}

func no_fiddle(t tom) {
    t.j = 15
}

func fiddle(t *tom) {
    t.j = 15
}

func (t *tom) play_with() {
    t.j = 12
}

func (t tom) no_play() {
    t.j = 12
}

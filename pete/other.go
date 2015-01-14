package main

import (
    "fmt"
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
    fmt.Println("Hello")
    f := tom{j: 3, f: fred{x: "yo", y: 5}}
    fmt.Println(f.f.x)
}

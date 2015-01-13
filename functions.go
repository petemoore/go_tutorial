package main

import (
    "fmt"
    "math"
	"code.google.com/p/go-tour/pic"
    "strings"
	"code.google.com/p/go-tour/wc"
//    "github.com/petemoore/go_tutorial"
)

var (
    c int
)

type IPAddr [4]byte

func add(x, y int) int {
    return x+y
}

func swap(x, y string) (A, B string) {
    A, B = y, x
    return
}

func fibonacci() func() int {
    last := 1
	penultimate :=0
	return func() int {
	    new_last := last + penultimate
	    penultimate = last
		last = new_last
		return last
	}
}

func main() {
    fmt.Println(add(42, 13))
    var (
        k bool = true
        a int = 3
    )
    ef, eg := "mary", "john"
    x, y := swap("hello", "pete")
    fmt.Println(x, y)
    fmt.Println(c, k, a, ef, eg)
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
	var X, Y int = 3, 4
	var F = math.Sqrt(float64(X*X + Y*Y))
	var Z int = int(F)
    fmt.Printf("%T %v\n", F, F)
    fmt.Printf("%T %v\n", Z, Z)
	fmt.Println(X, Y, Z)
	const (
        World = "世界"
        Pi = 3.14
    )
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println(Sqrt(77))
	wc.Test(WordCount)
	pic.Show(Pic)
	fibby := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibby())
	}
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
    main2()
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
	    return 0, ErrNegativeSqrt(x)
    }
	var oldz float64 = 0 
    z := 10.0
    for z-oldz > 1e-12 || oldz-z > 1e-12 {
        oldz=z
        z=z-(z*z-x)/(2*z)
    }   
    return z, nil
}

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := range image {
		image[y] = make([]uint8, dx)
		for x := range image[y] {
			image[y][x] = uint8(x^y)
		}
	}
	return image
}

func WordCount(s string) map[string]int {
    var entries []string = strings.Fields(s)
	totals := make(map[string]int)
	for _,word := range entries {
	    totals[word]++
	}
	return totals
}

func (x IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", x[0], x[1], x[2], x[3])
}

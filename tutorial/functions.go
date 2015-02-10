package main

import (
	"code.google.com/p/go-tour/pic"
	"code.google.com/p/go-tour/reader"
	"code.google.com/p/go-tour/tree"
	"code.google.com/p/go-tour/wc"
	"fmt"
	"github.com/petemoore/go_tutorial/lib"
	"io"
	"math"
	"os"
	"strings"
)

var (
	c int
)

type IPAddr [4]byte

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(new_bytes []byte) (int, error) {
	var old_bytes = make([]byte, len(new_bytes))
	read, err := reader.r.Read(old_bytes)
	if err != nil {
	}
	for i := range old_bytes {
		switch b := old_bytes[i]; {
		case (b >= 'a' && b <= 'm') || (b >= 'A' && b <= 'M'):
			new_bytes[i] = old_bytes[i] + 13
		case (b >= 'n' && b <= 'z') || (b >= 'N' && b <= 'Z'):
			new_bytes[i] = old_bytes[i] - 13
		default:
			new_bytes[i] = old_bytes[i]
		}
	}
	return read, err
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (A, B string) {
	A, B = y, x
	return
}

func fibonacci() func() int {
	penultimate, last := 0, 1
	return func() int {
		penultimate, last = last, last+penultimate
		return last
	}
}

func main() {
	fmt.Println(add(42, 13))
	var (
		k bool = true
		a int  = 3
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
		Pi    = 3.14
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
	if err := lib.Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	reader.Validate(MyReader{})
	str := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{str}
	io.Copy(os.Stdout, &r)
	fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
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
		oldz = z
		z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

type MyReader struct{}

func (reader MyReader) Read(bytes []byte) (int, error) {
	for b := range bytes {
		bytes[b] = byte(65)
	}
	return len(bytes), nil
}

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := range image {
		image[y] = make([]uint8, dx)
		for x := range image[y] {
			image[y][x] = uint8(x ^ y)
		}
	}
	return image
}

func WordCount(s string) map[string]int {
	var entries []string = strings.Fields(s)
	totals := make(map[string]int)
	for _, word := range entries {
		totals[word]++
	}
	return totals
}

func (x IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", x[0], x[1], x[2], x[3])
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

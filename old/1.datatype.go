package main

import (
	"fmt"
	"cmd/compile/internal/arm"
)

func main() {
	var n int = 10
	var d float64 = 10.45
	var b bool = true
	var s string = "อะไรก็ไม่รู้"
	var ns [3]int
	var bs []int
	var m map[string]string // ns[0]
	var pn *int

    // Primitive Datatype จะ Default ค่า Zero value ให้
	fmt.Println("int", n)
	fmt.Println("float64", d)
	fmt.Println("bool", b)
	fmt.Println("string", s)

	// แต่ Reference Datatype แบบนี้ต้องกำหนดค่าก่อนจึงจะใช้งานได้ เพราะเป็นแค่ป้ายบอกทาง ไม่ใช่กล่องเก็บของ
	fmt.Println("array int", ns)
	fmt.Println("slice byte", bs, bs == nil)
	fmt.Println("map[string]string", m, bs == nil)
	fmt.Println("pointer", pn, pn == nil)

    pn = new(int)
    *pn = 10
    fmt.Println("pointer", pn, pn == nil)

    m = map[string]string
    // m = make(map[string]string) // หรือแบบนี้ ซึ่งต่างจาก Slice คือไม่ต้องกำหนดขนาด
    m["name"] = "Kasem"
    m["sername"] = "Arnontavilas"
	//m["name"] = "satit"

    fmt.Println(m)

    bs = make([]int, 1, 2) // ต่างจาก map ตรงที่การ make ต้องกำหนดขนาด ในพารามิเตอร์ที่ 2 (type, len, cap)
    bs = append(bs, 20)
    bs = append(bs, 20)
    bs = append(bs, 20)
    bs = append(bs, 20)
    bs = append(bs, 20)
    bs = append(bs, 20)
    bs = append(bs, 20)
    bs = append(bs, 20)
    fmt.Println(bs)
    fmt.Println("len:",len(bs))
    fmt.Println("cap:", cap(bs))

    x := []int{1,2,3}
    y := []int{4,5,6}
    z := []int{7,8,9}
    w := len(x) + len(y)+ len(z)
    c := make([]int, 0)

	for _, v := range x {
		c = append(c, v)
	}
    c = append(c, x...)
    c = append(c, y...)
    c = append(c, z...)
    fmt.Println(c)
}

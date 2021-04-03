package main

import "fmt"

func main() {
	a := []int{0}
	a = append(a, 0)
	b := a[:]
	fmt.Printf("aDir: %p, bDir: %p\n", &a, &b)
	a = append(a, 2)
	fmt.Println("a[2]", a[2])

	b = append(b, 1)
	fmt.Println("b[2]", b[2])
	fmt.Printf("aDir: %p, bDir: %p\n", &a, &b)

	c := []int{0, 0}
	c = append(c, 0)
	d := c[:]
	fmt.Printf("cDir: %p, dDir: %p\n", &c, &d)
	c = append(c, 2)
	fmt.Println("c[3]", c[3])

	d = append(d, 1)
	fmt.Println("d[3]", d[3])
	fmt.Printf("cDir: %p, dDir: %p\n", &c, &d)
}

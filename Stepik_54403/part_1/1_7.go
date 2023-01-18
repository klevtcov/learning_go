package m1_7

import (
	"fmt"
)

// Константы и iota

const (
	a  = iota + 1 // a == 1  (iota == 0)
	_             // iota + 1 пропуск (iota == 1)
	b             // iota + 1  = 3  (iota == 2)
	c             // iota + 1  = 4  (iota == 3)
	d  = c + 2    // d = 6  (iota == 4)
	t             // t = 6 - копируем предыдущую строчку (iota == 5)
	i             // i = 6 - копируем предыдущую строчку (iota == 6)
	i2 = iota + 2 // 7 + 2 = 9 (iota == 7)
)

func main() {
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("t", t)
	fmt.Println("i", i)
	fmt.Println("i2", i2)
}

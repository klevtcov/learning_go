package main

import "fmt"

func main() {

	fmt.Println("")

	// continue, break, swith case

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i) - 1 3 5 7 9
	// }

	flag := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag = false
			break
		} else if i == 1 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(flag)

	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("TGIF")
	case "Mon", "Thu", "Wed":
		fmt.Println("Boring")
	default:
		fmt.Println("default")
	}

	switch {
	case day == "Fri":
		fmt.Println("TGIF2")
		break
	}

	err := errors.New("asda sdas asd ")
	switch err {
	case errr:
		fmt.Println("  ")
	case erRR:
		ftm.Println(" asd")
	}

}

package main

import (
	"fmt"
	"io/ioutil"
)

func getFiles(path string) ([]string, error) {
	var files []string
	dirFiles, err := ioutil.ReadDir(fmt.Sprintf("%s", path))
	if err != nil {
		return nil, err
	}
	for _, i := range dirFiles {
		files = append(files, i.Name())
	}
	return files, nil
}

func test(path string) ([]string, error) {
	return getFiles(path)
}

// явно указываем, что функция вернёт значение переменной rog со значением int
func returnRog() (rog int) {
	rog = 123
	return
}

func main() {

	privet := func(x, y int) int {
		i := (x + y) * 2
		return i
	}

	t, _ := test(".")
	fmt.Println((t))
	f, _ := getFiles(".")
	fmt.Println(f)            // [main.go]
	fmt.Println(returnRog())  // 123
	fmt.Println(privet(1, 5)) // 12

}

package ma0042

import "fmt"

func RemoveDuplicates(input []string) []string {
	output := make([]string, 0)
	inputSet := make(map[string]struct{}, len(input))
	for _, v := range input {
		if _, ok := inputSet[v]; !ok {
			output = append(output, v)

		}
		inputSet[v] = struct{}{}
	}

	return output
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println(input)

	fmt.Println(RemoveDuplicates(input))
}

// Представьте, что на входе есть слайс строк:
// input := []string{
//     "cat",
//     "dog",
//     "bird",
//     "dog",
//     "parrot",
//     "cat",
// }
// Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса:
// func RemoveDuplicates(input []string) []string

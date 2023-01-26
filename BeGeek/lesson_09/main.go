package main

import (
	"fmt"
	"sort"
)

func main() {

	// srezFloat64 := []float64{4.2, -1.3, 0.8, -3.8, 12.6}
	// fmt.Println(sort.Float64sAreSorted(srezFloat64))   // false
	// sort.Float64s(srezFloat64)                         // сортируем
	// fmt.Println(srezFloat64)                           // [-3.8 -1.3 0.8 4.2 12.6]
	// fmt.Println(sort.Float64sAreSorted(srezFloat64))   // true
	// fmt.Println(sort.SearchFloat64s(srezFloat64, 0.8)) // 2

	// --- STRINGS
	// s := []string{"one", "two", "tree"}
	// sort.Strings(s)
	// fmt.Println(s, sort.StringsAreSorted(s)) // [one tree two] true

	peopleInfo := []struct {
		Name string
		Age  int
	}{
		{"Be", 11},
		{"Geek", 65},
		{"You", 28},
		{"Tube", 75},
	}

	sort.Slice(peopleInfo, func(i, j int) bool { return peopleInfo[i].Name < peopleInfo[j].Name })
	fmt.Println(peopleInfo) // [{Be 11} {Geek 65} {Tube 75} {You 28}]
	sort.SliceStable(peopleInfo, func(i, j int) bool { return peopleInfo[i].Age < peopleInfo[j].Age })
	fmt.Println(peopleInfo) // [{Be 11} {You 28} {Geek 65} {Tube 75}]

}

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	r := csv.NewReader(rd)

	record, err := r.ReadAll()
	if err != nil && err != io.EOF {
		return err
	}

	if len(record) == 10 {
		fmt.Println(record[4][2])
	}

	return nil
}

func main() {
	const root = "./task" // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

}

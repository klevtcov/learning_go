package m_task

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

// Другие решения
//
//
// func main() {
// 	z, _ := zip.OpenReader("task.zip")
// 	defer z.Close()
// 	for _, f := range z.File {
// 		r, _ := f.Open()
// 		if rows, _ := csv.NewReader(r).ReadAll(); len(rows) == 10 && len(rows[4]) == 10 {
// 			fmt.Println(f.Name, rows[4][2])
// 		}
// 		r.Close()
// 	}
// }
//------------------
//
// func main() {
//     r, _ := zip.OpenReader("task.zip") //открываем файл архива
//     defer r.Close()                    // после выполнения всего закрываем файл архива

//     for _, file := range r.File { //перебираем все файл
//         if !file.FileInfo().IsDir() { // если файл есть, то
//             ff, _ := file.Open() // открываем каждый файл в цикле
//             if data, _ := csv.NewReader(ff).ReadAll(); len(data) == 10 { //проверка текущего файла в цикле, если ок то ---->
//                 fmt.Println(data[4][2]) // печатаем значения
//             }
//             ff.Close() // закрываем текущий файл в цикле и переходим к следующему
//         }
//     }
// }
//
// ----------------------
//
//
// func walkFunc(path string, info fs.FileInfo, err error) error{
// 	if err != nil {
// 	   return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
// 	}

// 	file, _ := os.Open(path)
// 	defer file.Close()

// 	r := csv.NewReader(file)
// 	records, err := r.ReadAll()
// 	if len(records) > 1 {
// 		fmt.Println(records[4][2])
// 		fmt.Println(info.Name())
// 		fmt.Println(path)
// 	}
// 	return nil
// }

// func main() {
// 	filepath.Walk("./task", walkFunc)
// }

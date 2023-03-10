package model

import "fmt"

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES($1, $2)", name, todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO WHERE name =$1", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

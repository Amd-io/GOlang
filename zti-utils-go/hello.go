package main

import (
	"./scanner"
	"fmt"
	"strconv"
)

type person struct {
	Id int
	Name string
	Surname string
}

func main() {
	var people []person
	scanner.CommandPrefix = "/"
	scanner.ParamPrefix = "-"
	scanner.Commands = append(scanner.Commands, scanner.Command{
		Name: "add",
		Desc: "Add new person",
		Consume: func(params map[string][]string) {
			if len(params["i"])==0 {fmt.Println("[ERROR] Add -i <id> param"); return }
			id, err := strconv.Atoi(params["i"][0])
			if err != nil {fmt.Println("[ERROR] -i <id> param must be a number!"); return }
			if len(params["n"])==0 {fmt.Println("[ERROR] Add -n <name> param"); return }
			if len(params["s"])==0 {fmt.Println("[ERROR] Add -s <surname> param"); return }
			people = append(people, person{Name: params["n"][0], Surname: params["s"][0], Id: id})
		},
	})
	scanner.Commands = append(scanner.Commands, scanner.Command{
		Name: "print",
		Desc: "Print all people",
		Consume: func(params map[string][]string) {
			for _, person := range people {
				fmt.Println(person)
			}
		},
	})
	scanner.Commands = append(scanner.Commands, scanner.Command{
		Name: "exit",
		Desc: "Stop the program",
		Consume: func(params map[string][]string) {
			scanner.Running = false
		},
	})
	scanner.Commands = append(scanner.Commands, scanner.Command{
		Name: "remove",
		Desc: "Delete person by given id",
		Consume: func(params map[string][]string) {
			if len(params["i"])==0 {fmt.Println("[ERROR] Add -i <id> param"); return }
			id, err := strconv.Atoi(params["i"][0])
			if err != nil {fmt.Println("[ERROR] -i <id> param must be a number!"); return }
			for index, candidate := range people{
				if candidate.Id != id {continue}
				people[index] = people[len(people)-1]
				people = people[:index+copy(people[index:], people[index+1:])]
			}
		},
	})
	scanner.Commands = append(scanner.Commands, scanner.Command{
		Name: "edit",
		Desc: "Edit person by given id",
		Consume: func(params map[string][]string) {
			if len(params["i"])==0 {fmt.Println("[ERROR] Add -i <id> param"); return }
			id, err := strconv.Atoi(params["i"][0])
			if err != nil {fmt.Println("[ERROR] -i <id> param must be a number!"); return }
			var person *person
			for index, candidate := range people{
				if candidate.Id != id {continue}
				person = &people[index]
			}
			if len(params["n"])!=0 {(*person).Name=params["n"][0]}
			if len(params["s"])!=0 {(*person).Surname=params["s"][0]}
		},
	})
	scanner.Start()
}

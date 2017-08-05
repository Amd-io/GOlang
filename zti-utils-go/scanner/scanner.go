package scanner

import (
	"bufio"
	"os"
	"strings"
	"sync"
	"fmt"
)

type Command struct {
	Name string
	Desc string
	Consume func(params map[string][]string)
}

var CommandPrefix string
var ParamPrefix string
var Running bool
var Commands []Command
var DefaultPrefix string = "DEFAULT"
var Welcome string = "\nThis is a default welcome message\nTo change is set scanner.Welcome string\nType " +
	"*CommandPrefix*help to get list of all commands"

func Scan()  {
	input := read()
	var consumers []Command
	//find command
	{
		var chars []string = strings.Split(input, "")
		chars = chars[strings.LastIndex(input, CommandPrefix)+1:]
		input = ""
		for _, c := range chars {
			input += c
		}
	}
	//find consumers
	{
		var name string
		var args []string = deleteEmpty(strings.Split(input, " "))
		if len(args) < 1 {
			fmt.Println("[ERROR] No such command!")
			return
		}
		name = args[0]
		args = args[1:]
		for _, consumer := range Commands {
			if strings.ToUpper(consumer.Name) != strings.ToUpper(name) {
				continue
			}
			consumers = append(consumers, consumer)
		}
		if len(consumers) < 1 {
			fmt.Println("[ERROR] No such command!")
			return
		}
		input = ""
		for _, token := range args {
			input += token+" "
		}
		input = strings.Trim(input, " ")
	}
	var params map[string][]string = make(map[string][]string)
	//create params map
	{
		var prefix string = DefaultPrefix
		var tokens []string
		for _, val := range deleteEmpty(strings.Split(input, " ")) {
			if strings.ContainsAny(val, ParamPrefix) {
				params[prefix] = tokens
				prefix = strings.Replace(val, "-", "", 1)
				tokens = make([]string, 0)
				continue
			}
			tokens = append(tokens, val)
		}
		params[prefix] = tokens
	}
	//pass params to all commands
	for _, command := range consumers {
		command.Consume(params)
	}

}

func Start()  {
	var wg sync.WaitGroup
	fmt.Println(Welcome)
	Commands = append(Commands, Command{
		Name: "help",
		Desc: "Print all available commands",
		Consume: func(params map[string][]string) {
			fmt.Println("AVAILABLE COMMANDS:")
			for _, command := range Commands {
				fmt.Println("\t"+CommandPrefix+command.Name+" - "+command.Desc)
			}
		},
	})
	go func() {
		Running = true
		for ;Running; {
			Scan()
		}
		defer wg.Done()
	}()
	wg.Add(1)
	wg.Wait()
}

func read() string {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = text[:len(text)-1]
	return text
}

func deleteEmpty(slice []string) []string {
	var result []string
	for _, element := range slice{
		if element == "" {continue}
		element = strings.Trim(element, " ")
		result = append(result, element)
	}
	return result
}
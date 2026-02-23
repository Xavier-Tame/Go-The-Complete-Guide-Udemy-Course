package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs-practice/note"
	"example.com/structs-practice/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo content: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	printSomething(todo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer value: ", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float value: ", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String value: ", stringVal)
		return
	}

	/*switch value.(type) {
	case int:
		fmt.Println("Integer value: ", value)
	case float64:
		fmt.Println("Float value: ", value)
	case string:
		fmt.Println("String value: ", value)
	default:
		fmt.Println("Unknown type")
	}*/
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data: ", err)
		return err
	}

	fmt.Println("Data saved successfully!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // For Windows compatibility

	return text
}

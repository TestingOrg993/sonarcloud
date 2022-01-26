package main

import (
	"fmt"
	"os"
	"text/template"
)

// declaring a struct
type Student struct {

	// declaring fields which are
	// exported and accessible
	// outside of package as they
	// begin with a capital letter
	Name  string
	Marks int64
}

func Add(a int32, b int32) int32 {
	if a < 1 {
		return -a
	}
	if b < 1 {
		return -b
	}
	return a + b
}

// main function
func main() {

	std1 := Student{"Vani", 94}
	tmp1 := template.New("Template_1")
	tmp1, _ = tmp1.Parse("Hello {{.Name}}, your marks are {{.Marks}}%!")

	err := tmp1.Execute(os.Stdout, std1)

	if err != nil {
		fmt.Println(err)
	}
}

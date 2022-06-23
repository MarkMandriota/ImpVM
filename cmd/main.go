package main

import (
	"fmt"
	"io"
	"log"
	"os"

	imp "github.com/MarkMandriota/ImpVM"
)

func main() {
	name := ""

	if len(os.Args) >= 2 {
		name = os.Args[1]
	} else {
		fmt.Printf("enter binary file name: ")
		fmt.Scan(&name)
	}

	m := imp.NewMachine(os.Stdout, os.Stdin)

	fs := tryAssign(os.Open(name))
	m.Text = tryAssign(io.ReadAll(fs))
	fs.Close()

	m.Run()
}

func tryAssign[T any](v T, e error) T {
	if e != nil {
		log.Fatalln(e)
	}

	return v
}

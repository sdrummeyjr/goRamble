package main

import (
	"fmt"
	"goRamble/evaluator"
	"goRamble/lexer"
	"goRamble/object"
	"goRamble/parser"
	"goRamble/repl"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func runProgram(filename string) {

	if filepath.Ext(filename) != ".rmbl" {
		fmt.Println("Can only parse and eval '.rmbl' files. Exiting the interpreter.")
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	f, err := ioutil.ReadFile(wd + "/" + filename)
	if err != nil {
		fmt.Println("ramble: ", err.Error())
		os.Exit(1)
	}
	l := lexer.New(string(f))
	p := parser.New(l, wd)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Println(p.Errors()[0])
		os.Exit(1)
	}
	scope := object.NewEnvironment()
	e := evaluator.Eval(program, scope)
	if e.Inspect() != "null" {
		fmt.Println(e.Inspect())
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Ramble programming language (goRamble implementation)!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else {
		runProgram(args[0])
	}

}

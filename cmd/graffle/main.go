package main

import (
	"github.com/PlagaMedicum/graffle/pkg/gen"
	"github.com/PlagaMedicum/graffle/pkg/listener"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"os"
	"os/exec"
)

const targetFName = "./.tmp/target.go"

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Provide files to parse!")
	}

	fname := os.Args[1]
	fs, err := antlr.NewFileStream(fname)
	if err != nil {
		log.Fatalf("Cannot read file \"%s\"", fname)
	}
	lexer := parser.NewGraffleLexer(fs)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGraffleParser(stream)

	s := &listener.GraffleListener{}
	antlr.ParseTreeWalkerDefault.Walk(s, p.File())

	err = os.Mkdir(".tmp", 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Error creating .tmp folder: %v", err)
	}

	f, err := os.OpenFile(targetFName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatalf("Error opening target code file: %v", err)
	}

	_, err = f.Write(s.Buffer.Bytes())
	if err != nil {
		log.Fatalf("Error writing target code in file %s: %v", targetFName, err)
	}

	err = exec.Command("/bin/sh", "-c", "go build -o build/build ./.tmp/target.go").Run()
	if err != nil {
		log.Fatalf("Error building target code file: %v", err)
	}
}

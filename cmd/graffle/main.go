package main

import (
	"github.com/PlagaMedicum/graffle/pkg/gen"
	"github.com/PlagaMedicum/graffle/pkg/listener"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"os"
)


func main() {
	fname := os.Args[1]
	fs, err := antlr.NewFileStream(fname)
	if err != nil {
		log.Fatalf("Cannot read file \"%s\"", fname)
	}
	lexer := parser.NewGraffleLexer(fs)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGraffleParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&listener.GraffleListener{}, p.File())
}

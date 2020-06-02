BIN=./bin/graffle_compiler
MAIN=./cmd/graffle/main.go

build: $(BIN)
$(BIN):
	go build -o $(BIN) $(MAIN)

.PHONY: antlr4gen
antlr4gen:
	antlr4 -o ./pkg/gen -listener -visitor -Dlanguage=Go -lib . ./GraffleParser.g4 ./GraffleLexer.g4
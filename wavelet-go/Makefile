
GOCMD = go
GOBUILD = $(GOCMD) build

ANTLR = antlr4

BUILD_DIR = .build
EXE_NAME = $(BUILD_DIR)/exe/app

antlr:
	cd antlr
	$(ANTLR) -Dlanguage=Go QuarkLexer.g4 -o ../gen
	$(ANTLR) -Dlanguage=Go QuarkParser.g4 -o ../gen


build:
	go build -o $(EXE_NAME) src/main.go

run: build
	./$(EXE_NAME)


clean:
	rm -rf $(BUILD_DIR)
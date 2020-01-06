
GOCMD = go
GOBUILD = $(GOCMD) build

ANTLR = `java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'

BUILD_DIR = .build
EXE_NAME = $(BUILD_DIR)/exe/app

testantlr:
	which $(ANTLR)

antlr:
	$(ANTLR) -Dlanguage=Go antlr/QuarkLexer.g4
	$(ANTLR) -Dlanguage=Go antlr/QuarkParser.g4


build:
	go build -o $(EXE_NAME) src/main.go

run: build
	./$(EXE_NAME)


clean:
	rm -rf $(BUILD_DIR)

.PHONY: antlr build clean
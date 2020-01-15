package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

type ParserTestInputData struct {
	Cases []ParseTestInputCase `yaml:"cases"`
}

type ParseTestInputCase struct {
	Name string `yaml:"name"`
	Input string `yaml:"quark"`
	Expect string `yaml:"expect"`
}

type TestErrorListener struct {
	errorCount int
}

func (t *TestErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Printf("syntax error line %d:%d\n", line, column)
	t.errorCount += 1
}

func (t *TestErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	println("ambiguity error")
	t.errorCount += 1
}

func (t *TestErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (t *TestErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {

}

func testCheckError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("error %e\n", err)
	}
}

func TestParsePackage(t *testing.T) {
	testData, err := ioutil.ReadFile("parser_test_data.yml")
	testCheckError(err, t)

	inputData := ParserTestInputData{}
	err = yaml.Unmarshal(testData, &inputData)
	testCheckError(err, t)

	for _, testCase := range inputData.Cases {
		testname := testCase.Name
		t.Run(testname, func(t *testing.T) {
			parser := NewStringParser(testCase.Input)
			testErrors := TestErrorListener{}
			parser.AddErrorListener(&testErrors)
			parser.GetPackageAST()

			if testCase.Expect == "valid" && testErrors.errorCount > 0 {
				t.Errorf("expected no errors but got %d errors", testErrors.errorCount)
			} else if testCase.Expect == "invalid" && testErrors.errorCount == 0 {
				t.Errorf("expected parser error but got 0 errors")
			}
		})
	}
}




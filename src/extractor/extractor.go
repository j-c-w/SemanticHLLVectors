package extractor

import (
	"golang.org/x/tools/go/analysis"
)

var Extractor = &analysis.Analyzer {
	Name: "Int Function Extractor",
	Doc: "Extract int->int functions",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, err) {

}

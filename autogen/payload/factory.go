package payload

import (
	"github.com/waflab/waflab/test"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
)

type payloadConverter func(value string, payload *test.Input) error

var converterFactory = map[int]payloadConverter{
	parser.TkVarArgs:       addArg,
	parser.TkVarArgsNames:  addArgNames,
	parser.TkVarArgsGet:    addArg, // equivalent to VarArgs
	parser.TkVarFiles:      addFiles,
	parser.TkVarFilesNames: addFilesNames,
}

func AddVariable(v *parser.Variable, value string, payload *test.Input) error {
	if f, ok := converterFactory[v.Tk]; ok {
		err := f(value, payload)
		if err != nil {
			return err
		}
		return nil
	}
	panic("not supported operator!")
}

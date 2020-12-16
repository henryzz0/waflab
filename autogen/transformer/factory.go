package transformer

import "github.com/hsluoyz/modsecurity-go/seclang/parser"

type transformReverser func(variable string) string

var reverserFactory = map[int]transformReverser{
	parser.TkTransLowercase:    reverseLowercase,
	parser.TkTransUrlDecode:    reverseUrlDecode,
	parser.TkTransUrlDecodeUni: reverseUrlDecode,
}

func ReverseTransform(transformers []*parser.Trans, variable string) string {
	for i := len(transformers) - 1; i >= 0; i-- {
		if f, ok := reverserFactory[transformers[i].Tk]; ok {
			variable = f(variable)
		}
	}
	return variable
}

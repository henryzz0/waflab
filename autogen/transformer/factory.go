package transformer

import (
	"log"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
)

type transformReverser func(variable string) string

var reverserFactory = map[int]transformReverser{
	parser.TkTransBase64Decode:       reverseBase64Decode,
	parser.TkTransCompressWhitespace: reverseCompressWhiteSpace,
	parser.TkTransHexDecode:          reverseHexDecode,
	parser.TkTransLength:             reverseLength,
	parser.TkTransNormalisePath:      reverseNormalizePath,
	parser.TkTransNormalizePath:      reverseNormalizePath,
	parser.TkTransNormalisePathWin:   reverseNormalizePathWin,
	parser.TkTransNormalizePathWin:   reverseNormalizePathWin,
	parser.TkTransLowercase:          reverseLowercase,
	parser.TkTransUrlDecode:          reverseUrlDecode,
	parser.TkTransUrlDecodeUni:       reverseUrlDecode,
}

func ReverseTransform(transformers []*parser.Trans, variable string) string {
	for i := len(transformers) - 1; i >= 0; i-- {
		if f, ok := reverserFactory[transformers[i].Tk]; ok {
			variable = f(variable)
		} else {
			log.Fatalf("Unsupported transformation %s", parser.TransformationNameMap[transformers[i].Tk])
		}
	}
	return variable
}

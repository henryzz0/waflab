package operator

import "github.com/hsluoyz/modsecurity-go/seclang/parser"

type operationReverser func(argument string, not bool) (string, error)

var reverserFactory = map[int]operationReverser{
	parser.TkOpRx: reverseRx,
	parser.TkOpBeginsWith: reverseBeginsWith,
	parser.TkOpContains: reverseContains,
	parser.TkOpEndsWith: reverseEndWith,
}

func ReverseOperator(operator *parser.Operator) (string, error) {
	if operator.Not {
		panic("negative match not supported yet!")
	}
	if f, ok := reverserFactory[operator.Tk]; ok {
		return f(operator.Argument, operator.Not)
	}
	panic("not supported operator!")
}
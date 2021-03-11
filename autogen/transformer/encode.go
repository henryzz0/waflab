// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package transformer

import (
	"encoding/hex"
	"fmt"
)

func rune2HexString(r rune) string {
	return hex.EncodeToString([]byte{byte(r)})
}

// https://www.w3.org/TR/CSS2/syndata.html#characters
// Encode characters using CSS 2.x escape rules, Ex: '\000026' -> '&'
func cssEncode(r rune) string {
	return fmt.Sprintf("\\%x", r)
}

// https://www.w3.org/TR/REC-html40/charset.html#h-5.3
// The syntax "&#D;", where D is a decimal number, refers to the ISO 10646 decimal character number D
func htmlDecimalEncode(r rune) string {
	return fmt.Sprintf("&#%03d;", r) // &#DDD decimal number
}

// https://www.w3.org/TR/REC-html40/charset.html#h-5.3
// The syntax "&#xH;" or "&#XH;", where H is a hexadecimal number,
// refers to the ISO 10646 hexadecimal character number H.
// Hexadecimal numbers in numeric character references are case-insensitive.
func htmlHexEncode(r rune) string {
	return fmt.Sprintf("&#x%2s;", rune2HexString(r)) // &#xHH, hexadecimal
}

func jsHexEncode(r rune) string {
	return fmt.Sprintf("\\u%04s", rune2HexString(r))
}

// http://www.ecma-international.org/ecma-262/6.0/#sec-names-and-keywords
// \OOO (where O is any octal number)
func jsOctalEncode(r rune) string {
	if int32(r) < 0 || int32(r) > 255 {
		return string(r)
	}
	return fmt.Sprintf("\\%03o", r)
}

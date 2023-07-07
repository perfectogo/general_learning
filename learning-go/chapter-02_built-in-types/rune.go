package chapter02builtintypes

import (
	"fmt"
)

func Rune() {
	var (
		char       rune = 'a'          // single unicode character
		octal_8    rune = '\323'       // 8-bit octal
		hex_8      rune = '\x21'       // 8-bit hecadecimal
		hex_16     rune = '\u0032'     // 16-bit hecadecimal
		unicode_32 rune = '\U00000061' // 32-bit unicode

		newLine     rune = '\n'
		newTab      rune = '\t'
		singleQuote rune = '\''
		backSlash   rune = '\\'
	)

	fmt.Println("// single unicode character:", char)
	fmt.Println("// 8-bit octal:", octal_8)
	fmt.Println("// 8-bit hecadecimal:", hex_8)
	fmt.Println("// 16-bit hecadecimal:", hex_16)
	fmt.Println("// 32-bit unicode:", unicode_32)
	fmt.Println("newLine:", (newLine))
	fmt.Println("newTab:", newTab)
	fmt.Println("singleQuote:", singleQuote)
	fmt.Println("backSlash:", backSlash)
}

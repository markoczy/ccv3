package parser

func IsOneOf(runes ...rune) func(rune) bool {
	return func(r rune) bool {
		for _, v := range runes {
			if r == v {
				return true
			}
		}
		return false
	}
}

func IsEqual(ref rune) func(rune) bool {
	return func(r rune) bool {
		return ref == r
	}
}

func IsNumeric(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return r == '.'
}

func IsStrictNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsIdentifier(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return r == '_'
}

var IsOperator = IsOneOf('+', '-', '/', '*', '^')
var IsControl = IsOneOf('(', ')')

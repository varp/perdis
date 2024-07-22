package parser

func isLetter(character byte) bool {
	if (character >= 'a' && character <= 'z') ||
		(character >= 'A' && character <= 'Z') ||
		(character >= '0' && character <= '9') ||
		(character == '_' || character == '*' || character == '/') {
		return true
	}

	return false
}

func isSpace(character byte) bool {
	return character == ' ' || character == '\t' || character == '\n'
}

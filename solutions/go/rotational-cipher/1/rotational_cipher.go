package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	result := ""

	for _, value := range plain {
		switch {
		case value >= 'A' && value <= 'Z':
			result += string('A' + (value-'A'+rune(shiftKey))%26)
		case value >= 'a' && value <= 'z':
			result += string('a' + (value-'a'+rune(shiftKey))%26)
		default:
			result += string(value)
		}
	}

	return result
}

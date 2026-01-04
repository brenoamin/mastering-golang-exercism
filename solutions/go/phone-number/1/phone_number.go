package phonenumber

import (
    "errors"
    "strings"
    "fmt"
)

func Number(phoneNumber string) (string, error) {
	return parse(phoneNumber)
}


func AreaCode(phoneNumber string) (string, error) {
	n, err := parse(phoneNumber)
	if err != nil {
		return "", err
	}
	return n[:3], nil
}


func Format(phoneNumber string) (string, error) {
	n, err := parse(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}

func parse(phoneNumber string) (string, error) {
	n := normalize(phoneNumber)

	if len(n) == 11 {
		if n[0] != '1' {
			return "", errors.New("invalid country code")
		}
		n = n[1:]
	}

	if len(n) != 10 {
		return "", errors.New("invalid length")
	}

	if n[0] <= '1' || n[3] <= '1' {
		return "", errors.New("invalid area or exchange code")
	}

	return n, nil
}

func normalize(input string) string {
	var b strings.Builder

	for _, r := range input {
		if r >= '0' && r <= '9' {
			b.WriteByte(byte(r))
		}
	}

	return b.String()
}


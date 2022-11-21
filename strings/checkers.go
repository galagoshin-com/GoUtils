package strings

import "strconv"

func IsNumeric(text string) bool {
	_, err := strconv.Atoi(text)
	_, err2 := strconv.ParseFloat(text, 16)
	return err == nil && err2 == nil
}

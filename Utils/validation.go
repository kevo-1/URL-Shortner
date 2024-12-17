package utils

func ValidateString(gvnHash string, str string) bool {
	MD5Hash(str)
	return gvnHash == MD5Hash(str)
}
package helper

import "strings"

func RetrieveID(name string) string {
	return strings.Split(name, "-")[0]
}

package a

import "log"

func Assert(expr bool, message ...interface{}) {
	if !expr {
		log.Fatalln("assertion is false", message)
	}
}

func AssertErr(err error) {
	if err != nil {
		panic(err)
	}
}

func UnwrapBytes(bytes []byte, err error) []byte {
	AssertErr(err)
	return bytes
}

func AssertEqual(expected string, value string) {
	if expected != value {
		panic(expected + " != " + value)
	}
}

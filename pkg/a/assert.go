package a

import "log"

func Assert(expr bool, message ...string) {
	if !expr {
		log.Println("assertion is false", message)
	}
}

func AssertErr(err error) {
	if err != nil {
		panic(err)
	}
}

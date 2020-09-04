package main

import (
	"strings"
	"testing"
)

func Test_Greeting(t *testing.T) {
	result := greeting("teste")

	test1 := strings.HasPrefix(result, "<b>")
	test2 := strings.HasSuffix(result, "</b>")

	if !test1 || !test2 {
		t.Errorf("test fail! no return <b> string </b>")
	}
}

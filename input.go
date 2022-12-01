package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadInput(day int) string {
	data, err := ioutil.ReadFile(fmt.Sprint(day) + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

package input

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func ReadInput(day int) string {
	data, err := ioutil.ReadFile(filepath.Join(fmt.Sprint(day), fmt.Sprint(day)+".txt"))
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

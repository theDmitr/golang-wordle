package main

import (
	"fmt"
	. "main/logic"
	. "main/logic/generator"
	. "main/logic/storage"
)

func main() {
	storage := NewStorage("123")
	generator := NewGenerator()
	wordle := NewWordle(storage, generator)

	var att string = "aeaea"
	result, err := wordle.Attempt(&att)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

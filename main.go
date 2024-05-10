package main

import (
	"main/logic"
	"main/logic/generator"
	"main/logic/storage"
	"main/ui/web"
)

func main() {
	storage := storage.NewStorage("123")
	generator := generator.NewGenerator()
	logic := logic.NewWordle(storage, generator)
	web := web.NewWordleUi(logic)
	web.Start()
}

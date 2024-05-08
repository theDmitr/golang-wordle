package generator

import (
	. "main/logic/storage"
	"math/rand"
)

type IGenerator interface {
	GenerateWord(storage IWordStorage) *string
}

type Generator struct{}

func NewGenerator() *Generator {
	generator := new(Generator)
	return generator
}

func (generator *Generator) GenerateWord(storage IWordStorage) *string {
	words := *storage.Words()
	return &words[rand.Intn(len(words))]
}

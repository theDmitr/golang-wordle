package logic

import (
	"errors"
	. "main/logic/generator"
	. "main/logic/storage"
	"strings"
)

const (
	letterStateRight = iota
	letterStateExists
	letterStateAbsent
)

type WordleGame struct {
	storage   IWordStorage
	generator IGenerator
	word      string
	attempts  byte
}

func NewWordle(storage IWordStorage, generator IGenerator) *WordleGame {
	wordle := new(WordleGame)
	wordle.storage = storage
	wordle.generator = generator
	wordle.Regenerate()
	return wordle
}

func (wordle *WordleGame) Attempt(attemptWord *string) (*[]int, error) {
	if wordle.attempts == 0 {
		return nil, errors.New("attempts are over")
	}

	if len(*attemptWord) != len(wordle.word) {
		return nil, errors.New("word length invalid")
	}

	result := make([]int, len(wordle.word))

	for i := 0; i < len(wordle.word); i++ {
		if wordle.word[i] == (*attemptWord)[i] {
			result[i] = letterStateRight
		} else if strings.Contains(wordle.word, string((*attemptWord)[i])) {
			result[i] = letterStateExists
		} else {
			result[i] = letterStateAbsent
		}
	}

	wordle.attempts--
	return &result, nil
}

func (wordle *WordleGame) Regenerate() {
	wordle.word = *wordle.generator.GenerateWord(wordle.storage)
	wordle.attempts = 6
}

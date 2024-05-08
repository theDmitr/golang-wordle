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
	word     string
	attempts byte
}

func NewWordle(storage IWordStorage, generator IGenerator) *WordleGame {
	wordle := new(WordleGame)
	wordle.word = *generator.GenerateWord(storage)
	wordle.attempts = 6
	return wordle
}

func (wordle *WordleGame) Attempt(attemptWord *string) (*map[int]int, error) {
	if wordle.attempts == 0 {
		return nil, errors.New("attempts are over")
	}

	if len(*attemptWord) != len(wordle.word) {
		return nil, errors.New("word length invalid")
	}

	resultMap := make(map[int]int)

	for i := 0; i < len(wordle.word); i++ {
		if wordle.word[i] == (*attemptWord)[i] {
			resultMap[i] = letterStateRight
		} else if strings.Contains(wordle.word, string((*attemptWord)[i])) {
			resultMap[i] = letterStateExists
		} else {
			resultMap[i] = letterStateAbsent
		}
	}

	wordle.attempts--
	return &resultMap, nil
}

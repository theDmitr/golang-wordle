package utils

import "os"

func LoadFile(filename *string) (*[]byte, error) {
	body, err := os.ReadFile(*filename)

	if err != nil {
		return nil, err
	}

	return &body, nil
}

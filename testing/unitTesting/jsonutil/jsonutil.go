package jsonutil

import (
	"encoding/json"
	"os"
)

type Human struct {
	Name string
	Age int
	Gender bool
}

func EncodeHuman2File(h *Human, filename string) error {
	dstFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func() {
		_ = dstFile.Close()
	}()
	encoder := json.NewEncoder(dstFile)
	err = encoder.Encode(h)
	if err != nil {
		return err
	}
	return nil
}

func DecodeJsonFile(filename string, receiver interface{}) error {
	srcFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = srcFile.Close()
	}()
	decoder := json.NewDecoder(srcFile)
	err = decoder.Decode(receiver)
	if err != nil {
		return err
	}
	return nil
}

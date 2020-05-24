package main

import (
	"fmt"

	"os"

	"github.com/vanhtuan0409/dbml-go/token"

	"github.com/vanhtuan0409/dbml-go/scanner"
)

func main() {
	f2, _ := os.Open("test.dbml")
	s := scanner.NewScanner(f2)
	for {
		l, c := s.LineInfo()
		tok, lit := s.Read()
		if tok == token.EOF {
			break
		}
		fmt.Printf("[%d:%d]\tToken: %s\tlit: %s\n", l, c, tok, lit)
	}
}

package main

import "fmt"

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Println("Storing into db", value)
	return nil
}

// This is coming from a third party lib
type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("FOO")
}

func myExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("my ex func", s)
		db.Store(s)
	}
}

func main() {
	s := &Store{}
	Execute(myExecuteFunc(s))
}

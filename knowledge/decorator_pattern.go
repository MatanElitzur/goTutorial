package knowledge

import (
	"fmt"
	"net/http"
)

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Println("Storing into db", value)
	return nil
}

//	func myExexuteFunc(s string) {
//		// access to DB
//		fmt.Println("my ex func", s)
//	}
func myExexuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("my ex func", s)
		db.Store(s)
	}
}

func makeHttpFunc(db DB, fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(db, w, r); err != nil {
			//
		}
	}
}

func handler(db DB, w http.ResponseWriter, r *http.Request) error {
	return nil
}

type httpFunc func(db DB, w http.ResponseWriter, r *http.Request) error

func DecoratorPattern_1() {
	s := &Store{}
	http.HandleFunc("/", makeHttpFunc(s, handler))
	execute(myExexuteFunc(s))
	//execute(myExexuteFunc)
}

// This is comming from a third party lib that we want to decorate
type ExecuteFn func(string)

func execute(fn ExecuteFn) {
	fn("I love Golang")
}

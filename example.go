package testlib

import (
	"github.com/vmihailenco/msgpack/v4"
)

type Person struct {
	Name string `msgpack:"name,omitempty"`
	Age  int32  `msgpack:"age,omitempty"`
}

func GetPerson(name string, age int32) ([]byte, error) {
	return msgpack.Marshal(&Person{
		Name: name,
		Age: age,
	})
}

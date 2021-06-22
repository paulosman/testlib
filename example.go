package testlib

type Person struct {
	Name string
	Age  int32
}

func GetPerson(name string, age int32) *Person {
	return &Person{
		Name: name,
		Age: age,
	}
}

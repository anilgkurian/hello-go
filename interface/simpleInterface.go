package main

type Person interface {
	name() string
}

type Employee struct {
	empName string
}

func (t Employee) name() string {
	return "Employee: " + t.empName
}

func main() {

	var e Person = Employee{"Anil"}
	println(e.name())

}

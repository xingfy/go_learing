package encap

import "testing"

type Employee struct {
    Id   string
    Name string
    Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
    e := Employee{"0", "Bob", 20}
    e1 := Employee{Name: "Bob", Age: 20}
    e2 := new(Employee)
    e2.Id = "2"
    e2.Name = "Rose"
    e2.Age = 22
    t.Log(e)
    t.Log(e1)
    t.Log(e2)
    t.Logf("e is %T", e)
    t.Logf("e2 is %T", e2)
}

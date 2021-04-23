package extension_test

import (
    "fmt"
    "testing"
)

type Pet struct {
}

func (p *Pet) Speck() {
    fmt.Println("...")
}

func (p *Pet) SpeckTo(name string) {
    p.Speck()
    fmt.Println(name)
}

type Dog struct {
    p *Pet
}

func (d *Dog) Speck() {
   fmt.Println("Wang!")
}

func (d *Dog) SpeckTo(name string) {
   d.Speck()
   fmt.Println(name)
}

func TestDog(t *testing.T) {
    dog := new(Dog)
    dog.SpeckTo("aaa")
}

type Cat struct {
    Pet
}

func (d *Cat) Speck() {
    fmt.Println("Miao!")
}

func TestCat(t *testing.T) {
    cat := new(Cat)
    cat.SpeckTo("aaa")
}

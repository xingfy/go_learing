package interface_test

import "testing"

type Program interface {
    WriteHelloWorld() string
}

type GoProgram struct {
}

func (g *GoProgram) WriteHelloWorld() string {
    return "fmt.Println(\"Hello world\")"
}

func TestClient(t *testing.T) {
    var p Program
    p = new(GoProgram)
    t.Log(p.WriteHelloWorld())
}

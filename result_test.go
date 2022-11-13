package result_test

import (
	"testing"
	"github.com/rawleyfowler/result"
)

func TestOk(t *testing.T) {
	o := result.Ok[string, any]("Hello")

	if !o.IsOk() {
		t.Fatal("Expected OK but got Error")
	}

	if o.IsError() {
		t.Fatal("Expected OK but got Error")
	}
}

func TestError(t *testing.T) {
	o := result.Error[any, string]("Hello")

	if o.IsOk() {
		t.Fatal("Expected Error but got OK")
	}

	if !o.IsError() {
		t.Fatal("Expected Error but got OK")
	}
}

func TestUnwrapDefault(t *testing.T) {
	o := result.Error[any, string]("Uhh ohh ERROR!!")

	j := o.UnwrapOr("Default")

	if j != "Default" {
		t.Fatalf("Expected default unwrap to be %s, was %s", "Default", j)
	}
}

func TestBind(t *testing.T) {
	o := result.Ok[string, any]("Yes!")

	// Would be nice if we could use >>= or (.) chaining :(
	m := result.Bind(o,
		func (a string) *result.Result[string, any] {
			if a == "Yes!" {
				return result.Ok[string, any]("We did it!")
			} else {
				return result.Ok[string, any]("We didn't do it!")
			}
		},
	)

	if m.Unwrap() != "We did it!" {
		t.Fatal("Binding failed to return new result")
	}
}

func TestMap(t *testing.T) {
	o := result.Ok[string, any]("Hello")

	p := result.Map(o, func (k string) int {
		return 2
	})

	if !p.IsOk() {
		t.Fatal("Exp[ected OK but got Error")
	}

	g := p.Unwrap()
	
	if g != 2 {
		t.Fatalf("Expected unwrap to yield %d but got %d", 2, g)
	}
}

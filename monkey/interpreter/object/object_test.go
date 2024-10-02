package object

import (
	"testing"
)

func TestIntegerInspect(t *testing.T) {
	integer := &Integer{Value: 123}
	if integer.Inspect() != "123" {
		t.Errorf("Inspect() = %s; want 123", integer.Inspect())
	}

	if integer.Type() != INTEGER_OBJ {
		t.Errorf("Type() = %s; want INTEGER_OBJ", integer.Type())
	}
}

func TestBooleanInspect(t *testing.T) {
	boolean := &Boolean{Value: true}
	if boolean.Inspect() != "true" {
		t.Errorf("Inspect() = %s; want true", boolean.Inspect())
	}

	if boolean.Type() != BOOLEAN_OBJ {
		t.Errorf("Type() = %s; want BOOLEAN_OBJ", boolean.Type())
	}
}

func TestNullInspect(t *testing.T) {
	null := &Null{}
	if null.Inspect() != "null" {
		t.Errorf("Inspect() = %s; want null", null.Inspect())
	}

	if null.Type() != NULL_OBJ {
		t.Errorf("Type() = %s; want NULL_OBJ", null.Type())
	}
}

func TestReturnValueInspect(t *testing.T) {
	rv := &ReturnValue{Value: &Integer{Value: 123}}
	if rv.Inspect() != "123" {
		t.Errorf("Inspect() = %s; want 123", rv.Inspect())
	}

	if rv.Type() != RETURN_VALUE_OBJ {
		t.Errorf("Type() = %s; want RETURN_VALUE_OBJ", rv.Type())
	}
}

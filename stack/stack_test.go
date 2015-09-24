package stack

import "testing"

func TestPush(t *testing.T) {
	stack := &Stack{}
	stack.Push(TValue(1))

	if len(stack.data) != 1 {
		t.Errorf("Expected data to have length 1, but have %d", len(stack.data))
	}

	if stack.data[0] != TValue(1) {
		t.Errorf("Expect first element to be equal to 1, but it is %d", stack.data[0])
	}
}

func TestPushManyElements(t *testing.T) {
	stack := &Stack{}
	stack.Push(TValue(1))
	stack.Push(TValue(2))

	if len(stack.data) != 2 {
		t.Errorf("Expected data to have length 1, but have %d", len(stack.data))
	}

	if stack.data[0] != TValue(1) {
		t.Errorf("Expect first element to be equal to 1, but it is %d", stack.data[0])
	}

	if stack.data[1] != TValue(2) {
		t.Errorf("Expect second element to be equal to 2, but it is %d", stack.data[0])
	}
}

func TestPop(t *testing.T) {
	stack := &Stack{}
	stack.Push(TValue(1))
	stack.Push(TValue(2))

	value, err := stack.Pop()

	if err != nil {
		t.Error(err)
	}

	if *value != TValue(2) {
		t.Errorf("Expect retrieved element to be equal to 2, but it is %d", stack.data[0])
	}

	if len(stack.data) != 1 {
		t.Errorf("Expected data length to be 1, but it is %d", len(stack.data))
	}
}

func TestPopWhenStackIsEmpty(t *testing.T) {
	stack := &Stack{}
	value, err := stack.Pop()

	if err == nil {
		t.Error("Expected error to have occurred")
	}

	if value != nil {
		t.Errorf("Expected value to be nil, but it is %v", *value)
	}
}

package stack

import "testing"

var itemOne TValue = new(int)
var itemTwo TValue = new(int)

func TestPush(t *testing.T) {
	stack := &Stack{}
	stack.Push(itemOne)

	if len(stack.data) != 1 {
		t.Errorf("Expected data to have length 1, but have %d", len(stack.data))
	}

	if stack.data[0] != itemOne {
		t.Errorf("Expect first element to be equal to 1, but it is %d", stack.data[0])
	}
}

func TestPushManyElements(t *testing.T) {
	stack := &Stack{}
	stack.Push(itemOne)
	stack.Push(itemTwo)

	if len(stack.data) != 2 {
		t.Errorf("Expected data to have length 1, but have %d", len(stack.data))
	}

	if stack.data[0] != itemOne {
		t.Errorf("Expect first element to be equal to 1, but it is %d", stack.data[0])
	}

	if stack.data[1] != itemTwo {
		t.Errorf("Expect second element to be equal to 2, but it is %d", stack.data[0])
	}
}

func TestPop(t *testing.T) {
	stack := &Stack{}
	stack.Push(itemOne)
	stack.Push(itemTwo)

	value, err := stack.Pop()

	if err != nil {
		t.Error(err)
	}

	if value != itemTwo {
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
		t.Errorf("Expected value to be nil, but it is %v", value)
	}
}

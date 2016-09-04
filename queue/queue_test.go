package queue

import "testing"

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	if queue.count != 0 {
		t.Fatal("queue count should equal zero")
	}
	if len(queue.v) != 0 {
		t.Fatal("queue length should equal zero")
	}
}

func TestQueuePush(t *testing.T) {
	queue := NewQueue()
	queue.Push("test")
	if queue.count == 0 {
		t.Fatal("queue count should not equal zero")
	}
	if queue.count != 1 {
		t.Fatal("queue count should equal one")
	}
	if len(queue.v) == 0 {
		t.Fatal("queue length should not equal zero")
	}

	if len(queue.v) != 1 {
		t.Fatal("queue length should equal one")
	}
}

func TestQueuePop(t *testing.T) {
	queue := NewQueue()
	data := queue.Pop()
	if data != nil {
		t.Fatal("queue data should be nil")
	}

	queue.Push("data")
	data = queue.Pop()
	if data == nil {
		t.Fatal("queue data should not be nil")
	}
	if queue.count != 0 {
		t.Fatal("queue length should equal zero")
	}

}

func TestQueueCount(t *testing.T) {
	queue := NewQueue()
	if queue.Count() != 0 {
		t.Fatal("queue count should be zero")
	}

	queue.Push("data")
	queue.Push("data")
	if queue.Count() != 2 {
		t.Fatal("queue count should be 2")
	}
	queue.Pop()
	if queue.Count() != 1 {
		t.Fatal("queue count should be 1")
	}

}

func TestQueuePushPop(t *testing.T) {
	queue := NewQueue()
	queue.Push("data1")
	queue.Push("data2")
	queue.Push("data3")
	data := queue.Pop()

	if data != "data1" {
		t.Fatal("data should equal 'data1'")
	}

	if queue.count != 2 {
		t.Fatal("queue count should be 2")
	}
	if queue.v[0] != "data2" {
		t.Fatal("First item of queue should be 'data2'")
	}

}

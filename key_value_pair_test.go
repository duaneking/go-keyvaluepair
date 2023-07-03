package kvp

import (
	"testing"
)

const test string = "test"
const key string = "key"
const value string = "value"

func TestCreateIntStringKeyValuePair(t *testing.T) {
	pair := KeyValuePair[int, string]{Key: 1, Value: test}

	if pair.Key != 1 {
		t.Errorf("Expected pair.Key to be 1, got %d", pair.Key)
	}

	if pair.Value != test {
		t.Errorf("Expected 'test', got %s", pair.Value)
	}
}

func TestChangeFloat64tStringKeyValuePair(t *testing.T) {
	pair := KeyValuePair[float64, string]{Key: 1.1, Value: test}

	if pair.Key != 1.1 {
		t.Errorf("Expected pair.Key of 1.1, got %v", pair.Key)
	}

	if pair.Value != test {
		t.Errorf("Expected pair.Value of 'test', got %s", pair.Value)
	}

	pair.Key = 3.2
	pair.Value = "test3"

	if pair.Key != 3.2 {
		t.Errorf("Expected Key value of 3, got %v", pair.Key)
	}

	if pair.Value != "test3" {
		t.Errorf("When change KeyValuePair expect the Value to be \"test2\", got %s", pair.Value)
	}
}

func TestDeconstructKeyValuePair(t *testing.T) {
	pair := KeyValuePair[string, string]{Key: key, Value: value}

	k, v := pair.Deconstruct()

	if k != key {
		t.Errorf("expected key, got %s", key)
	}

	if v != value {
		t.Errorf("On deconstruct expected value, got %s", pair.Value)
	}
}


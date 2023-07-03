package kvp

var Version string = "1.0.0.0"

type KeyValuePair[K any, V any] struct {
	Key   K `json:"Key" form:"Key"`
	Value V `json:"Value" form:"Value"`
}

func (pair KeyValuePair[TKey, TValue]) Deconstruct() (TKey, TValue) {
	return pair.Key, pair.Value
}

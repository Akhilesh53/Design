package main

type KeyVal struct {
	key   string
	value string
}

// NewKeyVal creates a new KeyVal with the given key and value.
func NewKeyVal(key string, value string) *KeyVal {
	return &KeyVal{
		key:   key,
		value: value,
	}
}

// write all getter and setter for keyval
func (kv *KeyVal) GetKey() string {
	return kv.key
}

func (kv *KeyVal) SetKey(key string) {
	kv.key = key
}

func (kv *KeyVal) GetValue() string {
	return kv.value
}

func (kv *KeyVal) SetValue(value string) {
	kv.value = value
}

func (kv *KeyVal) GetKeyVal() (string, string) {
	return kv.key, kv.value
}

func (kv *KeyVal) SetKeyVal(key string, value string) {
	kv.key = key
	kv.value = value
}

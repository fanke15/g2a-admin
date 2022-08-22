package basic

import iter "github.com/json-iterator/go"

var json = iter.ConfigCompatibleWithStandardLibrary

func Marshal(data interface{}) []byte {
	v, _ := json.Marshal(data)
	return v
}

func UnMarshal(data []byte, v interface{}) {
	_ = json.Unmarshal(data, v)
}

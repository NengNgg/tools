package json

import (
	jsoniter "github.com/json-iterator/go"
	"tools/pkg/bytesconv"
)

var jsonL = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal2String[T any](v T) (string, error) {
	bytes, err := jsonL.Marshal(v)
	if err != nil {
		return "", err
	}
	return bytesconv.BytesToString(bytes), nil
}

func Marshal2Bytes[T any](v T) ([]byte, error) {
	bytes, err := jsonL.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
func UnmarshalFromBytes[T any](data []byte) (T, error) {
	var v T
	err := jsonL.Unmarshal(data, &v)
	return v, err
}

func UnmarshalFromString[T any](data string) (T, error) {
	return UnmarshalFromBytes[T](bytesconv.StringToBytes(data))
}

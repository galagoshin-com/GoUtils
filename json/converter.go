package json

import (
	"encoding/json"
)

type Json string

func Encode(data map[string]any) (Json, error) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return Json(jsonString), nil
}

func Decode(input Json) (map[string]any, error) {
	r := []byte(input)
	var dat map[string]any
	err := json.Unmarshal(r, &dat)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func ConverInt(value any) int {
	return int(value.(float64))
}

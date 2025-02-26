package utils

import "encoding/json"

func MapToStruct(data map[string]interface{}, target interface{}) error {
	jsonBytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonBytes, target); err != nil {
		return err
	}

	return nil
}

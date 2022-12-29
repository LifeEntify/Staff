package staff

import (
	"encoding/json"
)

func NewStaff() Staff {
	return Staff{}
}
func (s *Staff) ToJson() (any, error) {
	resultByte, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var result any
	err = json.Unmarshal(resultByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

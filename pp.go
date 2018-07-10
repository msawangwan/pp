package pp

import "encoding/json"

func Print(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func PrintSafe(i interface{}) (string, error) {
	s, e := json.MarshalIndent(i, "", "\t")
	return string(s), e
}

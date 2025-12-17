package helloworld

import "encoding/json"

// innerFunctionForPanic - just return a string in bytes and error conversion in case there is one
func innerFunctionForPanic(anything string) ([]byte, error) {
	return json.Marshal(anything)
}

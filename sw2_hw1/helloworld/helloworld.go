package helloworld

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HelloWorldHandler - handler for /helloworld endpoint
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello World!\n"}
	// convert helloWorldResponse struct into json bytes
	// data - []byte, err - in case conversion failed
	data, err := json.Marshal(response)
	if err != nil {
		// in this case panic never occurs
		panic("Ooops")
	}
	// end result - {"Message":"Hello World!\n"}
	fmt.Fprintf(w, string(data))
}

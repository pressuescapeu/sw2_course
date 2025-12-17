package helloworld

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GoodbyeWorldHandler - handler for /goodbyeworld endpoint
func GoodbyeWorldHandler(w http.ResponseWriter, r *http.Request) {
	// function is not serializable so won't convert into json
	data, err := json.Marshal(innerFunctionForPanic)
	if err != nil {
		// and this will cause panic
		// but other endpoints still work btw
		panic("Ooops")
	}

	fmt.Fprintf(w, string(data))
}

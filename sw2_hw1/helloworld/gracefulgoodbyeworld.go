package helloworld

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GracefulGoodbyeWorldHandler(w http.ResponseWriter, r *http.Request) {
	// function is not serializable so won't convert into json
	data, err := json.Marshal(innerFunctionForPanic)
	if err != nil {
		// but instead of panic, we just return 500 and a message below
		http.Error(w, "internal server error\nGoodbye World!\n", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, string(data))
}

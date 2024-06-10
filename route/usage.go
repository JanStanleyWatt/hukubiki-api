package route

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonUsage struct {
	Description string `json:"description"`
}

func Usage(w http.ResponseWriter, r *http.Request) {
	usage := jsonUsage{
		Description: "`hukubiki` is a random number generator.",
	}

	res, err := json.Marshal(usage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(res))
}

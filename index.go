package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/docker/docker/pkg/namesgenerator"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := namesgenerator.GetRandomName(0)
	separator, ok := r.URL.Query()["separator"]
	if ok {
		name = strings.ReplaceAll(name, "_", separator[0])
	}

	switch r.Header.Get("accept") {
	case "application/json":
		out, err := json.Marshal(map[string]string{"name": name})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, string(out))
	default:
		fmt.Fprintf(w, name)
	}
}

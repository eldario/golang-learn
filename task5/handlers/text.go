package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type textOptions struct {
	Number int
	Text   string
}

func (m *Mappa) Text(writer http.ResponseWriter, request *http.Request) {
	var t textOptions

	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	m.result.Store(t.Number, t.Text)
	fmt.Fprint(writer, t.Number, t.Text)
}

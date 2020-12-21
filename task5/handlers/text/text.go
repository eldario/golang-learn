package text

import (
	"encoding/json"
	"net/http"
	counters "tasks/task5/handlers"
)

type readLiner interface {
	Read(string, int)
}

type Handler struct {
	lineReader readLiner
	counter    *counters.Metric
}

type payload struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.counter.Calls.Inc()
	p := payload{}

	defer request.Body.Close()

	err := json.NewDecoder(request.Body).Decode(&p)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	h.lineReader.Read(p.Text, p.Number)

	type result struct {
		Status   string `json:"status"`
		Chapter  int    `json:"chapter"`
		TextSize int    `json:"text_size"`
	}
	var jsonData []byte

	jsonData, _ = json.Marshal(result{
		Status:   "ok",
		Chapter:  p.Number,
		TextSize: len(p.Text),
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonData)
}

func New(lineReader readLiner) *Handler {

	var counter = counters.New("text", "text_func_calls")

	return &Handler{lineReader: lineReader, counter: counter}
}

package metrics

import (
	"net/http"
)

//var metrics = &Metrics{
//	calls: promauto.NewCounter(prometheus.CounterOpts{Namespace: "task5", Name: "superjob_calls"}),
//}

type Handler struct{}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}

func New() *Handler {
	return &Handler{}
}

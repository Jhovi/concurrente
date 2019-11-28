package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TransaccionRoutes struct {
	MyRouter http.Handler
}

func (transaccion *TransaccionRoutes) Routes() {
	transaccion.MyRouter = mux.NewRouter()

}

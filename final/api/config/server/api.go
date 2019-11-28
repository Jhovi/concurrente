package server

import (
	"encoding/json"
	"net/http"

	"../../app/controller"
	"github.com/gorilla/mux"
)

type Server interface {
	Router() http.Handler
}

type Api struct {
	router http.Handler
}

func (api *Api) Router() http.Handler {
	return api.router
}

var pu string

func InitServer(puerto string) Server {
	api := &Api{}
	pu = puerto
	router := mux.NewRouter()

	transaccionController := controller.TransaccionController{}
	transaccionController.SetMyInfo(puerto)

	router.HandleFunc("/", LinkOfRutas).Methods("GET")

	transaccionRouter := router.PathPrefix("/app").Subrouter()

	transaccionRouter.HandleFunc("/transaccion", transaccionController.RegistrarTransaccion).Methods("POST")
	transaccionRouter.HandleFunc("/transacciones", transaccionController.ListarTransaccionAll).Methods("GET")

	transaccionRouter.HandleFunc("/unidos", transaccionController.TodaLaRed).Methods("GET")
	transaccionRouter.HandleFunc("/unirse", transaccionController.UnirseALaRed).Methods("POST")
	transaccionRouter.HandleFunc("/unidos", transaccionController.NotificandoLaRed).Methods("POST")
	transaccionRouter.HandleFunc("/recivir", transaccionController.RecivirBlockChain).Methods("POST")
	// transaccionRouter.HandleFunc("/myhost", transaccionController.MyInfoHost).Methods("GET")
	api.router = router
	return api
}

func LinkOfRutas(w http.ResponseWriter, r *http.Request) {

	result := "la ruta general esta en localhost:" + pu + "/app"
	json.NewEncoder(w).Encode(result)
}

package model

type Transaccion struct {
	Monto   string `json:"monto,omitempty"`
	Origen  string `json:"origen,omitempty"`
	Destino string `json:"destino,omitempty"`
}

package controller

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"../model"
)

var direccionesRemotas []model.Host
var transaccionesBD []model.Transaccion
var Puerto string
var Ip string

var localBlockChain BlockChain

// type Transaccion struct {
// 	Monto   string `json:"monto,omitempty"`
// 	Origen  string `json:"origen,omitempty"`
// 	Destino string `json:"destino,omitempty"`
// }

type Block struct {
	Index        int
	Timestamp    time.Time
	Data         model.Transaccion
	PreviousHash string
	Hash         string
}

type BlockChain struct {
	Chain []Block
}

type TransaccionController struct {
}

func (controller *TransaccionController) SetMyInfo(puerto string) {
	Puerto = puerto
	Ip = "localhost"

	myHost := model.Host{Ip: Ip, Puerto: Puerto}
	direccionesRemotas = append(direccionesRemotas, myHost)

	localBlockChain = CreateBlockChain()
}
func (blockChain *BlockChain) AddBlock(block Block) {
	block.Timestamp = time.Now()
	block.Index = blockChain.GetLatesBlock().Index + 1
	block.PreviousHash = blockChain.GetLatesBlock().Hash
	block.Hash = block.CalculateHash()
	blockChain.Chain = append(blockChain.Chain, block)
}
func (blockChain *BlockChain) GetLatesBlock() Block {
	n := len(blockChain.Chain)
	return blockChain.Chain[n-1]
}
func (block *Block) CalculateHash() string {
	src := fmt.Sprintf("%d-%s-%s", block.Index, block.Timestamp.String(), block.Data)
	return base64.StdEncoding.EncodeToString([]byte(src))
}
func CreateBlockChain() BlockChain {
	bc := BlockChain{}
	genesis := bc.CreateGenesisBlock()
	bc.Chain = append(bc.Chain, genesis)
	return bc
}
func (blockChain *BlockChain) CreateGenesisBlock() Block {
	block := Block{
		Index:        0,
		Timestamp:    time.Now(),
		Data:         model.Transaccion{},
		PreviousHash: "0",
	}
	block.Hash = block.CalculateHash()
	return block
}

func (controller *TransaccionController) MyInfoHost(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(Puerto)
}
func (controller *TransaccionController) TodaLaRed(response http.ResponseWriter, request *http.Request) {
	if direccionesRemotas == nil {
		vacio := []string{}
		dataJson, _ := json.Marshal(vacio)
		response.Header().Set("Content-Type", "application/json")
		response.Write(dataJson)
	} else {
		dataJson, _ := json.Marshal(direccionesRemotas)
		response.Header().Set("Content-Type", "application/json")
		response.Write(dataJson)
		fmt.Print("Los host en la red son: ", direccionesRemotas)
	}
	fmt.Println()
}
func (controller *TransaccionController) NotificandoLaRed(response http.ResponseWriter, request *http.Request) {
	var allHost []model.Host
	_ = json.NewDecoder(request.Body).Decode(&allHost)

	direccionesRemotas = allHost
	json.NewEncoder(response).Encode(direccionesRemotas)
	fmt.Print("Los host en la red son: ", direccionesRemotas)
	fmt.Println()

}

func (controller *TransaccionController) UnirseALaRed(response http.ResponseWriter, request *http.Request) {
	var host model.Host
	_ = json.NewDecoder(request.Body).Decode(&host)

	direccionesRemotas = append(direccionesRemotas, host)

	for _, host := range direccionesRemotas {
		if host.Puerto != Puerto {
			clienteHttp := &http.Client{}
			url := "http://" + host.Ip + ":" + host.Puerto + "/app/unidos"
			log.Printf("conectandose a '%s'", url)
			dataComoJson, _ := json.Marshal(direccionesRemotas)
			peticion, _ := http.NewRequest("POST", url, bytes.NewBuffer(dataComoJson))
			respuesta, _ := clienteHttp.Do(peticion)
			defer respuesta.Body.Close()
			cuerpoRespuesta, _ := ioutil.ReadAll(respuesta.Body)
			respuestaString := string(cuerpoRespuesta)
			log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
		}

	}

	json.NewEncoder(response).Encode(true)

	fmt.Print("Nodo Agregado Exitosamente: ", host)
	fmt.Println()
}
func (controller *TransaccionController) RecivirBlockChain(response http.ResponseWriter, request *http.Request) {
	var bc BlockChain
	_ = json.NewDecoder(request.Body).Decode(&bc)

	localBlockChain = bc

	fmt.Print("el block chain es ", localBlockChain)
	fmt.Println()
}

func (controller *TransaccionController) RegistrarTransaccion(response http.ResponseWriter, request *http.Request) {
	var transaccion model.Transaccion
	_ = json.NewDecoder(request.Body).Decode(&transaccion)

	newBlock := Block{
		Data: transaccion,
	}
	localBlockChain.AddBlock(newBlock)

	for _, host := range direccionesRemotas {
		if host.Puerto != Puerto {
			clienteHttp := &http.Client{}
			url := "http://" + host.Ip + ":" + host.Puerto + "/app/recivir"
			log.Printf("conectandose a '%s'", url)
			dataComoJson, _ := json.Marshal(localBlockChain)
			peticion, _ := http.NewRequest("POST", url, bytes.NewBuffer(dataComoJson))
			respuesta, _ := clienteHttp.Do(peticion)
			defer respuesta.Body.Close()
			cuerpoRespuesta, _ := ioutil.ReadAll(respuesta.Body)
			respuestaString := string(cuerpoRespuesta)
			log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
		}
	}

	//U_U transaccionesBD = append(transaccionesBD, transaccion)
	json.NewEncoder(response).Encode(true)

	fmt.Print("Transaccion Agregado Exitosamente: ", localBlockChain)
	fmt.Println()
}

func (controller *TransaccionController) ListarTransaccionAll(response http.ResponseWriter, request *http.Request) {

	// if localBlockChain == nil {
	// 	vacio := []string{}
	// 	// json.NewEncoder(response).Encode(vacio)
	// 	dataJson, _ := json.Marshal(vacio)
	// 	response.Header().Set("Content-Type", "application/json")
	// 	response.Write(dataJson)
	// } else {
	// 	blockes := localBlockChain.Chain[1:]
	// 	fmt.Print("blockes: ", blockes)

	// 	var dataRecord []model.Transaccion
	// 	for i, block := range blockes {
	// 		fmt.Printf("- - - Transaccion Record No. %d - - - \n", i+1)
	// 		dataRecord = append(dataRecord, block.Data)
	// 	}
	// 	// dataJson, _ := json.Marshal(transaccionesBD)
	// 	dataJson, _ := json.Marshal(dataRecord)
	// 	response.Header().Set("Content-Type", "application/json")
	// 	response.Write(dataJson)
	// 	// json.NewEncoder(response).Encode(transaccionesBD)
	// }

	blockes := localBlockChain.Chain[1:]
	fmt.Print("blockes: ", blockes)

	var dataRecord []model.Transaccion
	for i, block := range blockes {
		fmt.Printf("- - - Transaccion Record No. %d - - - \n", i+1)
		dataRecord = append(dataRecord, block.Data)
	}
	// dataJson, _ := json.Marshal(transaccionesBD)
	dataJson, _ := json.Marshal(dataRecord)
	response.Header().Set("Content-Type", "application/json")
	response.Write(dataJson)
	// json.NewEncoder(response).Encode(transaccionesBD)

	fmt.Print("transacciones: ", localBlockChain)
	fmt.Println()
}

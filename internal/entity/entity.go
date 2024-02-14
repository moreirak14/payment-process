package entity

type OrderPaymentRequest struct {
	viagemExternoId    int     `json:"viagemExternoId"`
	tipo               string  `json:"tipo"`
	formaPagamento     string  `json:"formaPagamento"`
	cpfCnpjContratado  string  `json:"cpfCnpjContratado"`
	nomeContratado     string  `json:"nomeContratado"`
	chavePix           string  `json:"chavePix"`
	cpfMotorista       string  `json:"cpfMotorista"`
	nomeMotorista      string  `json:"nomeMotorista"`
	valor              float64 `json:"valor"`
	tipoBanco          string  `json:"tipoBanco"`
	pagamentoExternoId int     `json:"pagamentoExternoId"`
	ibgeOrigem         int     `json:"ibgeOrigem"`
	ibgeDestino        int     `json:"ibgeDestino"`
}

func NewOrderRequest(
	viagemExternoId int,
	tipo string,
	formaPagamento string,
	cpfCnpjContratado string,
	nomeContratado string,
	chavePix string,
	cpfMotorista string,
	nomeMotorista string,
	valor float64,
	tipoBanco string,
	pagamentoExternoId int,
	ibgeOrigem int,
	ibgeDestino int,
) *OrderPaymentRequest {
	return &OrderPaymentRequest{
		viagemExternoId:    viagemExternoId,
		tipo:               tipo,
		formaPagamento:     formaPagamento,
		cpfCnpjContratado:  cpfCnpjContratado,
		nomeContratado:     nomeContratado,
		chavePix:           chavePix,
		cpfMotorista:       cpfMotorista,
		nomeMotorista:      nomeMotorista,
		valor:              valor,
		tipoBanco:          tipoBanco,
		pagamentoExternoId: pagamentoExternoId,
		ibgeOrigem:         ibgeOrigem,
		ibgeDestino:        ibgeDestino,
	}
}

func (order *OrderPaymentRequest) Process() (*OrderPaymentResponse, error) {
	// This is the request to the BBC payment
	//URL := "https://api.payments/bbc/"
	//response, err := http.Post(URL, "application/json", bytes.NewBuffer(data))
	//if err != nil {
	//	slog.Error(err.Error())
	//}
	//defer response.Body.Close()
	//
	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	slog.Error(err.Error())
	//}

	// This is the return of the BBC payment request
	orderResponse := NewOrderPaymentResponse("Pagamento efetuado com sucesso!", "success")
	return orderResponse, nil
}

type OrderPaymentResponse struct {
	Mensagem string `json:"Mensagem"`
	Status   string `json:"status"`
}

func NewOrderPaymentResponse(mensagem string, status string) *OrderPaymentResponse {
	return &OrderPaymentResponse{
		Mensagem: mensagem,
		Status:   status,
	}
}

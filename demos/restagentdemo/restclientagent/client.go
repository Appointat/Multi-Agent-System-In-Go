package restclientagent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	rad "gitlab.utc.fr/lagruesy/ia04/demos/restagentdemo"
)

type RestClientAgent struct {
	id       string
	url      string
	operator string
	arg1     int
	arg2     int
}

func NewRestClientAgent(id string, url string, op string, arg1 int, arg2 int) *RestClientAgent {
	return &RestClientAgent{id, url, op, arg1, arg2}
}

func (rca *RestClientAgent) treatResponse(r *http.Response) int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp rad.Response
	json.Unmarshal(buf.Bytes(), &resp)

	return resp.Result
}

func (rca *RestClientAgent) doRequest() (res int, err error) {
	req := rad.Request{
		Operator: rca.operator,
		Args:     [2]int{rca.arg1, rca.arg2},
	}

	// sérialisation de la requête
	url := rca.url + "/calculator"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	res = rca.treatResponse(resp)

	return
}

func (rca *RestClientAgent) Start() {
	log.Printf("démarrage de %s", rca.id)
	res, err := rca.doRequest()

	if err != nil {
		log.Fatal(rca.id, "error:", err.Error())
	} else {
		log.Printf("[%s] %d %s %d = %d\n", rca.id, rca.arg1, rca.operator, rca.arg2, res)
	}
}

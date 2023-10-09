package main

import (
	"fmt"

	"gitlab.utc.fr/lagruesy/ia04/demos/restagentdemo/restclientagent"
)

func main() {
	ag := restclientagent.NewRestClientAgent("id1", "http://localhost:8080", "+", 11, 1)
	ag.Start()
	fmt.Scanln()
}

package main

import (
	"fmt"

	ras "gitlab.utc.fr/lagruesy/ia04/demos/restagentdemo/restserveragent"
)

func main() {
	server := ras.NewRestServerAgent(":8080")
	server.Start()
	fmt.Scanln()
}

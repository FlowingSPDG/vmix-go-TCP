package main

import (
	"log"

	"github.com/FlowingSPDG/vmix-go-tcp"
)

func main() {
	v, err := vmixtcp.New("localhost")
	if err != nil {
		panic(err)
	}
	defer v.Close()

	resp1, err := v.TALLY()
	if err != nil {
		panic(err)
	}
	log.Printf("TALLY RESPONSE1 : %s\n", resp1)

	resp1, resp2, err := v.XML()
	if err != nil {
		panic(err)
	}
	log.Printf("XML RESPONSE1 : %s\n", resp1)
	log.Printf("XML RESPONSE2 : %s\n", resp2)

	resp1, err = v.FUNCTION("PreviewInput Input=1")
	if err != nil {
		panic(err)
	}
	log.Printf("FUNCTION RESPONSE : %s\n", resp1)

	if err := v.QUIT(); err != nil {
		panic(err)
	}
}

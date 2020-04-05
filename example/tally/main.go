package main

import (
	"github.com/FlowingSPDG/vmix-go-tcp"
	"log"
	"time"
)

func main() {
	v, err := vmixtcp.New("localhost")
	if err != nil {
		panic(err)
	}
	defer v.Close()

	resp, err := v.SUBSCRIBE("TALLY")
	if err != nil {
		panic(err)
	}
	log.Printf("SUBSCRIBE TALLY RESPONSE : %s\n", resp)
	time.Sleep(time.Second * 30)
}

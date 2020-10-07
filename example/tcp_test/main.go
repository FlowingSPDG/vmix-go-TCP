package main

import (
	"log"

	vmixtcp "github.com/FlowingSPDG/vmix-go-TCP"
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

	// If you want to parse XML, Comment-out following code to parse it
	// https://github.com/FlowingSPDG/vmix-go/blob/master/models.go#L12-L45
	//
	/*
		import {
			"encoding/xml"
			"github.com/FlowingSPDG/vmix-go"
		}
		v := vmixgo.Vmix{}
		err = xml.Unmarshal(body, &v)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal XML... %v", err)
		}
	*/

	resp1, err = v.FUNCTION("PreviewInput Input=1")
	if err != nil {
		panic(err)
	}
	log.Printf("FUNCTION RESPONSE : %s\n", resp1)

	if err := v.QUIT(); err != nil {
		panic(err)
	}
}

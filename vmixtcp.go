package vmixtcp

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// Vmix main object
type Vmix struct {
	Conn       *net.Conn
	subscribe  *net.Conn
	subhandler func(string)
}

func New(dest string) (*Vmix, error) {
	vmix := &Vmix{}
	c, err := net.Dial("tcp", dest+":8099")
	if err != nil {
		return nil, err
	}

	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), "\r\n", "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return nil, fmt.Errorf("Unknown ERR : %v", Resps[2:])
	}

	log.Printf("vMix TCP API Initialized... : %s\n", Resp)

	vmix.Conn = &c

	// SUBSCRIBE related...
	subscriber, err := net.Dial("tcp", dest+":8099")
	if err != nil {
		return nil, err
	}
	vmix.subscribe = &subscriber

	/*
		go func() {
			reader := bufio.NewReader(subscriber)
			for {
				data, err := reader.ReadString('\n')
				if err != nil {
					log.Printf("Unknown error on subscriber : %v\n", err)
					break
				}
				log.Printf("SUBSCRIBER DATA : %v\n", string(data))
			}
		}()
	*/

	return vmix, nil
}

func (v *Vmix) Close() {
	c := *v.Conn
	c.Close()

	sub := *v.subscribe
	sub.Close()
}

func (v *Vmix) XML() (string, string, error) {
	c := *v.Conn
	_, err := c.Write([]byte("XML\r\n"))
	if err != nil {
		return "", "", err
	}
	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), "\r\n", "")
	Resps := strings.Split(Resp, " ")

	BodyLen, err := strconv.Atoi(Resps[1])
	if err != nil {
		return "", "", err
	}

	BodyBuffer := make([]byte, BodyLen)
	BodyLength, _ := c.Read(BodyBuffer)
	Body := strings.ReplaceAll(string(BodyBuffer[:BodyLength]), "\r\n", "")

	return Resp, Body, nil
}

func (v *Vmix) Tally() (string, error) {
	c := *v.Conn
	_, err := c.Write([]byte("TALLY\r\n"))
	if err != nil {
		return "", err
	}
	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1011) // Maximum possible length = 9 + 1000 + 2 = 1011 bytes
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), "\r\n", "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[4:])
	}

	return Resp, nil
}

func (v *Vmix) FUNCTION(funcname string) (string, error) {
	c := *v.Conn
	_, err := c.Write([]byte(fmt.Sprintf("FUNCTION %s\r\n", funcname)))
	if err != nil {
		return "", err
	}
	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), "\r\n", "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}

	return Resp, nil
}

func (v *Vmix) SUBSCRIBE(command string) (string, error) {
	c := *v.subscribe
	_, err := c.Write([]byte(fmt.Sprintf("SUBSCRIBE %s\r\n", command)))
	if err != nil {
		return "", err
	}
	// c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), "\r\n", "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}
	v.subscribe = &c

	return Resp, nil
}

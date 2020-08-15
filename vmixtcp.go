package vmixtcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	// Terminate letter
	Terminate = "\r\n"
)

// Vmix main object
type Vmix struct {
	Conn       *net.Conn
	subscribe  *net.Conn
	subhandler func(string)
}

// New vmix instance
func New(dest string) (*Vmix, error) {
	vmix := &Vmix{}
	c, err := net.Dial("tcp", dest+":8099")
	if err != nil {
		return nil, err
	}

	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
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

	go func() {
		reader := bufio.NewReader(subscriber)
		for {
			data, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Unknown error on subscriber : %v\n", err)
				continue
			}
			log.Printf("SUBSCRIBER DATA : %v\n", string(data))
		}
	}()

	return vmix, nil
}

// Close connection
func (v *Vmix) Close() {
	c := *v.Conn
	c.Close()

	sub := *v.subscribe
	sub.Close()
}

// XML Gets XML data. Same as HTTP API.
func (v *Vmix) XML() (string, string, error) {
	c := *v.Conn
	_, err := c.Write([]byte("XML" + Terminate))
	if err != nil {
		return "", "", err
	}
	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	BodyLen, err := strconv.Atoi(Resps[1])
	if err != nil {
		return "", "", err
	}

	BodyBuffer := make([]byte, BodyLen)
	BodyLength, _ := c.Read(BodyBuffer)
	Body := strings.ReplaceAll(string(BodyBuffer[:BodyLength]), Terminate, "")

	return Resp, Body, nil
}

// TALLY Get tally status
func (v *Vmix) TALLY() (string, error) {
	c := *v.Conn
	_, err := c.Write([]byte("TALLY" + Terminate))
	if err != nil {
		return "", err
	}
	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1011) // Maximum possible length = 9 + 1000 + 2 = 1011 bytes
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[4:])
	}

	return Resp, nil
}

// FUNCTION Send function
func (v *Vmix) FUNCTION(funcname string) (string, error) {
	c := *v.Conn
	_, err := c.Write([]byte(fmt.Sprintf("FUNCTION %s%s", funcname, Terminate)))
	if err != nil {
		return "", err
	}
	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}

	return Resp, nil
}

// SUBSCRIBE Event
func (v *Vmix) SUBSCRIBE(command string) (string, error) {
	c := *v.subscribe
	_, err := c.Write([]byte(fmt.Sprintf("SUBSCRIBE %s%s", command, Terminate)))
	if err != nil {
		return "", err
	}
	// c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}
	v.subscribe = &c

	return Resp, nil
}

// UNSUBSCRIBE from event.
func (v *Vmix) UNSUBSCRIBE(command string) (string, error) {
	c := *v.subscribe
	_, err := c.Write([]byte(fmt.Sprintf("UNSUBSCRIBE %s%s", command, Terminate)))
	if err != nil {
		return "", err
	}
	// c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}
	v.subscribe = &c

	return Resp, nil
}

// QUIT Sends QUIT sigal
func (v *Vmix) QUIT() error {
	c := *v.Conn
	_, err := c.Write([]byte(fmt.Sprintf("QUIT %s", Terminate)))
	if err != nil {
		return err
	}
	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	// check slice length
	if Resps[1] != "OK" {
		return fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}

	return nil
}

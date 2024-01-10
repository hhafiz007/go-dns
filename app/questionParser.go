package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type DNSQuestion struct {
	Name  []string
	Type  uint16
	Class uint16
}

func (q *DNSQuestion) createQuestion() []byte {

	messageBuffer := new(bytes.Buffer)

	for _, label := range q.Name {
		hexLength, _ := hex.DecodeString(fmt.Sprintf("%02X", len(label))) // first we need length of label and then label in hex

		binary.Write(messageBuffer, binary.BigEndian, hexLength)

		for _, char := range label {
			hexString, _ := hex.DecodeString(fmt.Sprintf("%02X", char))
			// fmt.Println(hexString, char)
			binary.Write(messageBuffer, binary.BigEndian, hexString)
		}

	}
	binary.Write(messageBuffer, binary.BigEndian, uint8(0)) // termination by a null byte

	binary.Write(messageBuffer, binary.BigEndian, q.Type)
	binary.Write(messageBuffer, binary.BigEndian, q.Class)

	question := messageBuffer.Bytes()

	return question

}

func NewDNSQuestion() *DNSQuestion {
	return &DNSQuestion{
		Name:  []string{"google", "com"},
		Type:  1,
		Class: 1,
	}
}

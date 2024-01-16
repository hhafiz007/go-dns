package main

import (
	"bytes"
	"encoding/binary"
)

type DNSQuestion struct {
	Name  []byte
	Type  uint16
	Class uint16
}

func (q *DNSQuestion) createQuestion() []byte {

	messageBuffer := new(bytes.Buffer)

	// termination by a null byte
	binary.Write(messageBuffer, binary.BigEndian, q.Name)
	binary.Write(messageBuffer, binary.BigEndian, q.Type)
	binary.Write(messageBuffer, binary.BigEndian, q.Class)

	question := messageBuffer.Bytes()

	return question

}

func NewDNSQuestion() *DNSQuestion {
	return &DNSQuestion{
		Name:  []byte("\x0ccodecrafters\x02io\x00"),
		Type:  1,
		Class: 1,
	}
}

func DynamicDNSQuestion(buf []byte) *DNSQuestion {

	i := 96

	for {
		if buf[i] == 0 {
			break
		}
		i++
	}

	return &DNSQuestion{
		Name:  buf[96 : i+1],
		Type:  binary.BigEndian.Uint16(buf[i+1 : i+3]),
		Class: binary.BigEndian.Uint16(buf[i+3 : i+5]),
	}
}

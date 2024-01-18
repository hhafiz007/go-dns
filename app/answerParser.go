package main

import (
	"bytes"
	"encoding/binary"
)

type DNSAnswer struct {
	Name   []byte
	Type   uint16
	Class  uint16
	TTL    uint32
	Length uint16
	Data   []byte
}

func (a *DNSAnswer) createAnswer() []byte {

	messageBuffer := new(bytes.Buffer)

	// termination by a null byte
	binary.Write(messageBuffer, binary.BigEndian, a.Name)
	binary.Write(messageBuffer, binary.BigEndian, a.Type)
	binary.Write(messageBuffer, binary.BigEndian, a.Class)
	binary.Write(messageBuffer, binary.BigEndian, a.TTL)
	binary.Write(messageBuffer, binary.BigEndian, a.Length)
	binary.Write(messageBuffer, binary.BigEndian, uint32(a.Class))

	answer := messageBuffer.Bytes()

	return answer

}

func NewDNSAnswer() *DNSAnswer {
	return &DNSAnswer{
		Name:   []byte("\x0ccodecrafters\x02io\x00"),
		Type:   1,
		Class:  1,
		TTL:    60,
		Length: 4,
		Data:   []byte("\x08\x08\x08\x08"),
	}
}

func DynamicDNSAnswer(q *DNSQuestion) *DNSAnswer {

	return &DNSAnswer{
		Name:   q.Name,
		Type:   1,
		Class:  1,
		TTL:    60,
		Length: 4,
		Data:   []byte("\x08\x08\x08\x08"),
	}
}

package main

import (
	"bytes"
	"encoding/binary"
)

type DNSMessage struct {
	Header     []byte
	Question   []byte
	Answer     []byte
	Authority  []byte
	Additional []byte
}

func (d *DNSMessage) createMessage(buf []byte) []byte {

	header := createDynamicHeader(buf)
	header.Flags |= (1 << 15)
	// fmt.Println(header)
	hBytes := header.createHeader()
	question := NewDNSQuestion()
	qBytes := question.createQuestion()
	answer := NewDNSAnswer()
	aBytes := answer.createAnswer()
	var reply []byte

	reply = append(reply, hBytes...)
	reply = append(reply, qBytes...)
	reply = append(reply, aBytes...)

	return reply

}

type DNSHeader struct {
	ID uint16 // 16bits -> A random ID assigned to query packets. Response packets must reply with the same ID.

	// Flags contains the 16bit long DNS header flags
	// QR      uint8  // 1bit -> 1 for a reply packet, 0 for a question packet.
	// OPCODE  uint8  // 4bit -> Specifies the kind of query in a message.
	// AA      uint8  // 1bit -> 1 if the responding server "owns" the domain queried, i.e., it's authoritative.
	// TC      uint8  // 1bit -> 1 if the message is larger than 512 bytes. Always 0 in UDP responses.
	// RD      uint8  // 1bit -> Sender sets this to 1 if the server should recursively resolve this query, 0 otherwise.
	// RA      uint8  // 1bit -> Server sets this to 1 to indicate that recursion is available.
	// Z       uint8  // 3bit -> Used by DNSSEC queries. At inception, it was reserved for future use.
	// RCODE   uint8  // 4bit -> Response code indicating the status of the response.
	Flags uint16 // 16bits -> Flags

	QDCOUNT uint16 // 16bit -> Number of questions in the Question section.
	ANCOUNT uint16 // 16bit -> Number of records in the Answer section.
	NSCOUNT uint16 // 16bit -> Number of records in the Authority section.
	ARCOUNT uint16 // 16bit -> Number of records in the Additional section.
}

func (h *DNSHeader) createHeader() []byte {

	id := h.ID
	flags := h.Flags
	QDCOUNT := h.QDCOUNT
	ANCOUNT := h.ANCOUNT
	NSCOUNT := h.NSCOUNT
	ARCOUNT := h.ARCOUNT

	messageBuffer := new(bytes.Buffer)
	binary.Write(messageBuffer, binary.BigEndian, id)
	binary.Write(messageBuffer, binary.BigEndian, flags)
	binary.Write(messageBuffer, binary.BigEndian, QDCOUNT)
	binary.Write(messageBuffer, binary.BigEndian, ANCOUNT)
	binary.Write(messageBuffer, binary.BigEndian, NSCOUNT)
	binary.Write(messageBuffer, binary.BigEndian, ARCOUNT)

	header := messageBuffer.Bytes()

	return header

}

func NewDNSHeader() *DNSHeader {
	return &DNSHeader{
		ID:      1234,
		Flags:   0,
		QDCOUNT: 1,
		ANCOUNT: 1,
		NSCOUNT: 0,
		ARCOUNT: 0,
	}
}

// 1000000000000000

func createDynamicHeader(buf []byte) *DNSHeader {
	// buf[16] = 1
	// fmt.Println("buffer 17", buf)
	dnsHeader := &DNSHeader{
		ID:      binary.BigEndian.Uint16(buf[:2]),
		Flags:   binary.BigEndian.Uint16(buf[2:4]),
		QDCOUNT: binary.BigEndian.Uint16(buf[4:6]),
		ANCOUNT: binary.BigEndian.Uint16(buf[6:8]),
		NSCOUNT: binary.BigEndian.Uint16(buf[8:10]),
		ARCOUNT: binary.BigEndian.Uint16(buf[10:12]),
	}

	dnsHeader.QDCOUNT = 1
	return dnsHeader

}

func NewDNSMessage() *DNSMessage {
	return &DNSMessage{}
}

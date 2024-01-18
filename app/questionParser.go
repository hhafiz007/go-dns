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

func getQuestionsList(h *DNSHeader, buf []byte) []DNSQuestion {

	totalQuestions := h.QDCOUNT

	var questions []DNSQuestion

	start := 12

	for i := 0; i < int(totalQuestions); i++ {
		question, end := DynamicDNSQuestion(buf, start)
		questions = append(questions, *question)
		start = end
	}

	return questions

}

func DynamicDNSQuestion(buf []byte, start int) (*DNSQuestion, int) {

	i := start

	for {
		if buf[i] == 0 {
			break
		}
		i++
	}

	// var name []byte
	// fmt.Println("DynamicDNSQuestion start", start)

	// for buf[start] != 0 {
	// 	firstTwoBits := (buf[start] >> 6) & 0b11
	// 	offset := uint16(start)
	// 	fmt.Println("DynamicDNSQuestion offset", offset)
	// 	if firstTwoBits != 0 {

	// 		offset = uint16(buf[start]+buf[start+1]) & 0x3FFF
	// 		labelLength := buf[offset]
	// 		name = append(name, buf[offset:offset+uint16(labelLength)+1]...)
	// 		start += 2

	// 	} else {
	// 		labelLength := buf[offset]
	// 		name = append(name, buf[offset:offset+uint16(labelLength)+1]...)
	// 		start = int(offset + uint16(labelLength) + 1)

	// 	}

	// }

	// name = append(name, 0)

	// fmt.Println("index is", i, buf)

	return &DNSQuestion{
		Name:  buf[start : i+1],
		Type:  binary.BigEndian.Uint16(buf[i+1 : i+3]),
		Class: binary.BigEndian.Uint16(buf[i+3 : i+5]),
	}, i + 5
}

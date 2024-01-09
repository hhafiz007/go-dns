package main

type DNSMessage struct {
	Header     []byte
	Question   []byte
	Answer     []byte
	Authority  []byte
	Additional []byte
}

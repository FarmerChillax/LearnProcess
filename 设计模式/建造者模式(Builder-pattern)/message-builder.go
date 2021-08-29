package main

import (
	"sync"
)

type Message struct {
	Header *Header
	Body   *Body
}

type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}

type Body struct {
	Items []string
}

// Message对象的builder
type builder struct {
	once *sync.Once
	msg  *Message
}

// 返回builder
func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg:  &Message{Header: &Header{}, Body: &Body{}},
	}
}

// Message对象成员构建方法
func (b *builder) WithSrcAddr(srcAddr string) *builder {
	b.msg.Header.SrcAddr = srcAddr
	return b
}

func (b *builder) WithSrcPort(srcPort uint64) *builder {
	b.msg.Header.SrcPort = srcPort
	return b
}

func (b *builder) WithDestAddr(destAddr string) *builder {
	b.msg.Header.DestAddr = destAddr
	return b
}

func (b *builder) WithDestPort(destPort uint64) *builder {
	b.msg.Header.DestPort = destPort
	return b
}

func (b *builder) SetHeaderItem(key, value string) *builder {
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})
	b.msg.Header.Items[key] = value
	return b
}

func (b *builder) WithBodyItem(record string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}

func (b *builder) Build() *Message {
	return b.msg
}

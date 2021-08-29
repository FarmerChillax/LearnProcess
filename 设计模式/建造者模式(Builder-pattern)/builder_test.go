package main

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	message := Builder().
		WithSrcAddr("192.168.1.1").
		WithSrcPort(2333).
		WithDestAddr("192.168.1.2").
		WithDestPort(8080).
		SetHeaderItem("content-type", "application/json").
		WithBodyItem("record1").WithBodyItem("record2").
		Build()
	if message.Header.SrcAddr != "192.168.1.1" {
		t.Errorf("expect src address 192.168.1.1, but actual %s.", message.Header.SrcAddr)
	}
	if message.Body.Items[0] != "record1" {
		t.Errorf("expect body item0 record1, but actual %s.", message.Body.Items[0])
	}
}

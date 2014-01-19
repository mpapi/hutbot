package main

import (
	"bytes"
	"testing"
	"time"
)

func TestMessageEmpty(t *testing.T) {
	m := Message{&StreamMessager{}, "test", "#test", "contents", time.Now()}

	if m.Empty() {
		t.Error("Message.Empty true for non-empty message")
	}

	m.Contents = ""
	if !m.Empty() {
		t.Error("Message.Empty false for empty message")
	}
}

func TestResponseEmpty(t *testing.T) {
	r := Response{&CommandScript{}, nil, "contents", "", time.Now()}

	if r.Empty() {
		t.Error("Response.Empty true for non-empty response")
	}

	r.Contents = ""
	if !r.Empty() {
		t.Error("Response.Empty false for empty response")
	}
}

func TestStreamMessager(t *testing.T) {
	buf := new(bytes.Buffer)
	buf.WriteString("test message 1\n")
	buf.WriteString("test message 2\n")

	s := &StreamMessager{buf}

	responses := make(chan Response)
	messages := make(chan Message)

	go s.Process(messages, responses)

	msg := <-messages
	if messager, ok := msg.Messager.(*StreamMessager); !ok || messager != s {
		t.Error("StreamMessager doesn't set Messager to itself on its messages")
	}
	if msg.Sender != "stdin" {
		t.Error("StreamMessager should send messages with Sender == 'stdin'")
	}
	if msg.Channel != "stdin" {
		t.Error("StreamMessager should send messages with Channel == 'stdin'")
	}
	if msg.Contents != "test message 1" {
		t.Error("StreamMessager sent message with the wrong contents")
	}
}

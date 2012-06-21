package main

import (
	"../nsq"
	"github.com/bmizerany/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// ensure that we can push a message through a topic and get it out of a channel
func TestPutMessage(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(os.Stdout)

	topic := GetTopic("test_put_message", 10, os.TempDir())
	channel1 := topic.GetChannel("ch", 10, os.TempDir())

	msg := nsq.NewMessage(<-idChan, []byte("test"))
	topic.PutMessage(msg)

	outputMsg := <-channel1.ClientMessageChan
	assert.Equal(t, msg.Id, outputMsg.Id)
	assert.Equal(t, msg.Body, outputMsg.Body)
}

// ensure that both channels get the same message
func TestPutMessage2Chan(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(os.Stdout)

	topic := GetTopic("test_put_message_2chan", 10, os.TempDir())
	channel1 := topic.GetChannel("ch1", 10, os.TempDir())
	channel2 := topic.GetChannel("ch2", 10, os.TempDir())

	msg := nsq.NewMessage(<-idChan, []byte("test"))
	topic.PutMessage(msg)

	outputMsg1 := <-channel1.ClientMessageChan
	assert.Equal(t, msg.Id, outputMsg1.Id)
	assert.Equal(t, msg.Body, outputMsg1.Body)

	outputMsg2 := <-channel2.ClientMessageChan
	assert.Equal(t, msg.Id, outputMsg2.Id)
	assert.Equal(t, msg.Body, outputMsg2.Body)
}
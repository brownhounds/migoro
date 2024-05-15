package testing_shit

import (
	"bytes"
	"log"
	"migoro/dispatcher"
	"migoro/snapshots"
	"testing"
)

func TestShit(t *testing.T) {
	buffer := new(bytes.Buffer)
	log.SetFlags(0)
	log.SetOutput(buffer)
	dispatcher.Init()
	snapshots.ToMatchSnapshot(t, "Hello There!!!!!", "test1")
}

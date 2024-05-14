package testing_shit

import (
	"bytes"
	"fmt"
	"log"
	"migoro/dispatcher"
	"testing"
)

func TestShit(t *testing.T) {
	buffer := new(bytes.Buffer)
	log.SetFlags(0)
	log.SetOutput(buffer)
	dispatcher.Init()

	fmt.Print("dksijdihd" + buffer.String())
}

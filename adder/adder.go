package adder

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {
	num := NumMessage{
		Value: 57,
	}

	out, err := proto.Marshal(num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode message: %s", err.Error())
		os.Exit(1)
	}

	gotNum := NumMessage{}
	if err := proto.Unmarshal(out, gotNum); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode message: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("added value: %d\n", gotNum.Value)
}

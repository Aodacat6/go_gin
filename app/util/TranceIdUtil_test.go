package util

import (
	"fmt"
	"testing"
)

func Test_printTranceId(t *testing.T) {
	id := printTranceId()
	fmt.Println(id)
}

package blockbook

import (
	"fmt"
	"testing"
)

func TestGetStatus(t *testing.T) {
	svc := NewClient("https://ac-dev0.net:29136")
	status, err := svc.GetStatus()
	if err != nil {
		panic(err)
	}

	fmt.Println(status)
}

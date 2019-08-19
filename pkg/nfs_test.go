package pkg

import (
	"fmt"
	"nfs-api/pkg/sg"
	"testing"
)

func TestRun(t *testing.T) {
	api := sg.NewInstance("https", "192.168.3.60", "6080", "optadmin", "adminadmin")
	done, err := api.CreateDirectory("/test/4455")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(done)
}

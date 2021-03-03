package go_auto_release_test

import (
	go_auto_release "github.com/angeliski/go-auto-release"
	"testing"
)

func TestRun(t *testing.T) {
	err := go_auto_release.Run()

	if err != nil {
		t.Errorf("Fail to pass %v", err)
	}
}

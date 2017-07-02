package exposevars

import (
	"testing"
	"time"
)

func TestExpose(t *testing.T) {
	if err := Port(12345); err != nil {
		t.Error(err)
	}
	<-time.NewTimer(2 * time.Minute).C
}

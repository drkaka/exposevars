package exposevars

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestExpose(t *testing.T) {
	if err := Start(12345, "abc"); err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	var info map[string]interface{}
	resp, err := http.Get("http://localhost:12345/debug/vars")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		t.Fatal(err)
	}

	// check uptime
	if info["uptime"].(float64) != float64(1) {
		t.Error("uptime wrong")
	}

	// check goroutines
	if info["goroutines"].(float64) <= float64(0) {
		t.Error("goroutine wrong")
	}

	// check goroutines
	if info["service"].(string) != "abc" {
		t.Error("service name wrong")
	}
}

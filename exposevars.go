package exposevars

import (
	// to expose /debug/vars
	_ "expvar"

	"fmt"
	"net"
	"net/http"
)

// Port to expose
func Port(port uint16) error {
	sock, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return err
	}
	go func() {
		http.Serve(sock, nil)
	}()
	return nil
}

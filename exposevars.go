package exposevars

import (
	// to expose /debug/vars
	"expvar"
	"runtime"
	"time"

	"fmt"
	"net"
	"net/http"
)

// Port to expose with
func Port(port uint16) error {
	var startTime = time.Now().UTC()

	expvar.Publish("goroutines", expvar.Func(func() interface{} {
		// return the goroutine count
		return runtime.NumGoroutine()
	}))

	expvar.Publish("uptime", expvar.Func(expvar.Func(func() interface{} {
		// return the uptime seconds
		return uint64(time.Since(startTime)) / 1e9
	})))

	sock, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return err
	}
	go func() {
		http.Serve(sock, nil)
	}()
	return nil
}

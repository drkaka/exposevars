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

// Start to expose with given port and service name.
func Start(port uint16, service string) error {
	var startTime = time.Now().UTC()

	expvar.Publish("goroutines", expvar.Func(func() interface{} {
		// return the goroutine count
		return runtime.NumGoroutine()
	}))

	expvar.Publish("uptime", expvar.Func(func() interface{} {
		// return the uptime seconds
		return uint64(time.Since(startTime)) / 1e9
	}))

	expvar.Publish("service", expvar.Func(func() interface{} {
		// return the service name
		return service
	}))

	sock, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return err
	}
	go func() {
		http.Serve(sock, nil)
	}()
	return nil
}

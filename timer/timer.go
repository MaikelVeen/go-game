package timer

import (
	"fmt"
	"log/slog"
	"time"
)

// Timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
//	defer timer("sum")()
func Timer(name string) func() {
	start := time.Now()
	return func() {
		slog.Info(fmt.Sprintf("%s executed", name), "duration", time.Since(start))
	}
}

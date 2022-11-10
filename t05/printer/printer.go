package printer

import "runtime"

func Version() string {
	return runtime.Version()
}

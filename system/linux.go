package system

import "runtime"

func RequireLinux() {
	if runtime.GOOS != "linux" {
		panic("DevBox currently supports Linux only")
	}
}

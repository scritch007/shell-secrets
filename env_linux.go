package shellsecret

import (
	"fmt"
)

func printEnv(key string) {

	fmt.Printf("%s=%s;export %s", envKey, key, envKey)
}

func init() {
	ShellCmd = LinuxShell
}

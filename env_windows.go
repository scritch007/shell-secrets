package shellsecret

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-ps"
)

func printEnv(key string) {

	if ShellCmd == PowerShell {
		fmt.Printf("$env:%s = '%s'", envKey, key)
		return
	}
	fmt.Printf(`
You need to run the following commands in order to setup the environment

for CMD.exe
set %s=%s

for PowerShell
$env:%s = '%s'`, envKey, key, envKey, key)
}

func init() {
	process, err := ps.FindProcess(os.Getppid())
	if err != nil {
		fmt.Printf("failed to find parent process")
	}
	switch process.Executable() {
	case "powershell.exe", "pwsh.exe":
		ShellCmd = PowerShell
	case "cmd.exe":
		ShellCmd = Cmd
	default:
		panic(fmt.Errorf("environment not supported: %s", process.Executable()))
	}
}

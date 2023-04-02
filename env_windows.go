package shellsecret

import "fmt"

func printEnv(key string) {
	fmt.Printf(`
for CMD.exe
set %s=%s

for PowerShell
$env:%s = '%s'`, envKey, key)
}

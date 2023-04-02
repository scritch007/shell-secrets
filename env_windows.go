package shellsecret

import "fmt"

func printEnv(key string) {
	fmt.Printf("set %s=%s", envKey, key)
}

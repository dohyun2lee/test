package testlib

import "fmt"

var pop map[string]string

func init() {
	pop = make(map[string]string)
	pop["puth"] = "switch"
	pop["dog"] = "free"
}

func GetMusic(singer string) string{
	return pop[singer]
}

func getKeys() {
	for _, kv := range pop {
		fmt.Println(kv)
	}
}
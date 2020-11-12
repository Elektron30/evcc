package main

//go:generate esc -o server/assets.go -pkg server -modtime 1566640112 -ignore .DS_Store dist

import (
	"github.com/mark-sch/evcc/cmd"
)

func main() {
	cmd.Execute()
}

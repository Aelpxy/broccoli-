package main

import (
	"github.com/aelpxy/fresh/cmd"
	"github.com/aelpxy/fresh/storage"
)

func main() {
	storage.NewFreshStorageSystem("store")
	cmd.Execute()
}

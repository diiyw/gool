package main

import (
	"os/exec"
)

func goUrl(v string) {
	cmd := exec.Command("cmd", "/C", "start", v)
	_ = cmd.Run()
}

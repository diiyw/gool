package main

import "os/exec"

func goUrl(v string) {
	cmd := exec.Command("open", v)
	_ = cmd.Run()
}

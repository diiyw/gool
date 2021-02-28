package main

func goUrl(v string) {
	cmd := exec.Command("bash", "-c", "xdg-open", v)
	_ = cmd.Run()
}

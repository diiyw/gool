package gool

func goUrl(v string) {
	cmd := exec.Command("start", v)
	_ = cmd.Run()
}

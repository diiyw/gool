package main

import (
	"github.com/getlantern/systray"
	"gopkg.in/yaml.v2"
	"os"
)

func main() {
	systray.Run(onReady, onExit)
}

type menu struct {
	Title  string   `json:"title"`
	Type   string   `json:"type"`
	Values []string `json:"values"`
	Tip    string   `json:"tip"`
	Pid    string   `json:"pid"`
}

var menuConfig struct {
	Menus []menu `json:"menus"`
}

func onReady() {
	conf, err := os.ReadFile("gool.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(conf, &menuConfig); err != nil {
		panic(err)
	}
	systray.SetIcon(getIcon("gool"))
	systray.SetTooltip("超级好用自定义的工具集")
	for _, menu := range menuConfig.Menus {
		m := systray.AddMenuItem(menu.Title, menu.Tip)
		switch menu.Type {
		case "shell":
			go shell(m, menu)
		case "service":
			go service(m, menu)
		case "builtin":
			go builtin(m, menu)
		}
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出Gool")
	go func() {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
		}
	}()
}

func onExit() {

}

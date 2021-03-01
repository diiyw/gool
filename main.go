package main

import (
	"github.com/getlantern/systray"
	"gopkg.in/yaml.v2"
	"os"
)

type menu struct {
	Title  string   `yaml:"title"`
	Type   string   `yaml:"type"`
	Values []string `yaml:"values"`
	Tip    string   `yaml:"tip"`
	Pid    string   `yaml:"pid"`
	Status string   `yaml:"status"`
	Menus  []menu   `yaml:"menus"`
}

var menuConfig struct {
	Menus []menu `yaml:"menus"`
}

func main() {
	conf, err := os.ReadFile("gool.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(conf, &menuConfig); err != nil {
		panic(err)
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("gool"))
	systray.SetTooltip("超级好用自定义的工具集")

	var hook = func(item *systray.MenuItem, m menu) {
		switch m.Type {
		case "line":
		case "service":
			go service(item, m)
		case "copy":
			go clipboardCopy(item, m)
		default:
			go shell(item, m)
		}
	}

	for _, menu := range menuConfig.Menus {
		m := systray.AddMenuItem(menu.Title, menu.Tip)
		if menu.Status != "" {
			m.Disable()
		}
		if len(menu.Menus) > 0 {
			for _, subMenu := range menu.Menus {
				sub := m.AddSubMenuItem(subMenu.Title, subMenu.Tip)
				if subMenu.Status != "" {
					sub.Disable()
				}
				hook(sub, subMenu)
				continue
			}
		}
		hook(m, menu)
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

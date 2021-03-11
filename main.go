package main

import (
	"github.com/diiyw/gotray"
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
	gotray.Run(ready)
}

func ready() {
	gotray.SetIcon(getIcon("gool"))
	for _, menu := range menuConfig.Menus {
		m := gotray.NewMenu().SetTitle(menu.Title).SetTooltip(menu.Tip)
		if menu.Status != "" {
			m.Disable()
		}
		gotray.AddMenu(m)
		if len(menu.Menus) > 0 {
			for _, subMenu := range menu.Menus {
				sub := gotray.NewMenu().SetTitle(subMenu.Title).SetTooltip(subMenu.Tip)
				m.AddSubMenu(sub)
				if subMenu.Status != "" {
					sub.Disable()
				}
				continue
			}
		}
	}
	gotray.AddMenu(gotray.NewMenu().SetTitle("Exit").Click(func() {
		os.Exit(0)
	}))
}

package main

import (
	"github.com/diiyw/gotray"
	"gopkg.in/yaml.v2"
	"os"
)

type menu struct {
	Title   string   `yaml:"title"`
	Type    string   `yaml:"type"`
	Values  []string `yaml:"values"`
	Tooltip string   `yaml:"tip"`
	Pid     string   `yaml:"pid"`
	Status  string   `yaml:"status"`
	Menus   []menu   `yaml:"menus"`
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
	_ = gotray.SetIcon(getIcon("gool"))
	for _, menu := range menuConfig.Menus {
		m := gotray.NewMenu().SetTitle(menu.Title).SetTooltip(menu.Tooltip)
		if menu.Status != "" {
			m.Disable()
		}
		switch menu.Type {
		case "shell":
			shell(m, menu)
		case "service":
			service(m, menu)
		}
		if len(menu.Menus) > 0 {
			for _, child := range menu.Menus {
				sub := gotray.NewMenu().SetTitle(child.Title).SetTooltip(child.Tooltip)
				if child.Status != "" {
					sub.Disable()
				}
				m.AddSubMenu(sub)
			}
		}
		gotray.AddMenu(m)
	}
	exit := gotray.NewMenu().SetTitle("Exit").Click(func() {
		gotray.Exit()
	})
	gotray.AddMenu(exit)
}

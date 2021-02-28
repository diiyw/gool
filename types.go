package main

import (
	"github.com/getlantern/systray"
	"os/exec"
	"runtime"
	"time"
)

func shell(item *systray.MenuItem, menu menu) {
	for {
		select {
		case <-item.ClickedCh:
			for _, she := range menu.Values {
				execCmd(she)
			}
		}
	}
}

func builtin(item *systray.MenuItem, menu menu) {
	switch menu.Values[0] {
	case "copyTimestamp":
		copyTimestamp(item, menu)
	case "openLink":
		openLink(item, menu)
	}
}

func execCmd(arg string) []byte {
	var result []byte
	var osExec = "/bin/sh"
	if runtime.GOOS == "windows" {
		osExec = "cmd"
	}
	var ret []byte
	switch runtime.GOOS {
	case "windows":
		ret, _ = exec.Command(osExec, "-c", arg).Output()
	default:
		ret, _ = exec.Command(osExec, "-c", arg).Output()
	}
	result = append(result, ret...)
	return result
}

func service(item *systray.MenuItem, menu menu) {
	item.SetIcon(getIcon("stop"))
	status := item.AddSubMenuItem("未启动", "")
	start := item.AddSubMenuItem("启动服务", "")
	start.SetIcon(getIcon("start"))
	restart := item.AddSubMenuItem("重启", "")
	restart.Disable()
	stop := item.AddSubMenuItem("停止", "")
	stop.Disable()
	var f = func() {
		pid := getPid(menu.Pid)
		if pid != "" {
			status.SetTitle("已启动(Pid:" + pid + ")")
			item.SetIcon(getIcon("running"))
			restart.Disabled()
			stop.Disabled()
		} else {
			item.SetIcon(getIcon("stop"))
			status.SetTitle("未启动")
			start.SetTitle("启动服务")
		}
		status.Disable()
	}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		f()
	}
	for {
		select {
		case <-item.ClickedCh:
		case <-restart.ClickedCh:
			execCmd(menu.Values[2])
			f()
		case <-start.ClickedCh:
			execCmd(menu.Values[0])
			f()
		}
	}
}

func openLink(item *systray.MenuItem, menu menu) {
	for {
		select {
		case <-item.ClickedCh:
			go goUrl(menu.Values[1])
		}
	}
}

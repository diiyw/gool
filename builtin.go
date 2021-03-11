package main

import (
	"github.com/atotto/clipboard"
	"github.com/diiyw/gotray"
	"github.com/golang-module/carbon"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func shell(m *gotray.Menu, menu menu) {
	m.Click(func() {
		for _, she := range menu.Values {
			execCmd(she)
		}
	})
}

func clipboardCopy(m *gotray.Menu, menu menu) {
	m.Click(func() {
		switch menu.Values[0] {
		case "timestamp":
			_ = clipboard.WriteAll(strconv.FormatInt(carbon.Now().ToTimestamp(), 10))
		case "date":
			_ = clipboard.WriteAll(carbon.Now().ToDateString())
		case "datetime":
			_ = clipboard.WriteAll(carbon.Now().ToDateTimeString())
		case "yesterday_timestamp":
			_ = clipboard.WriteAll(strconv.FormatInt(carbon.Yesterday().ToTimestamp(), 10))
		case "yesterday_date":
			_ = clipboard.WriteAll(carbon.Yesterday().ToDateString())
		case "yesterday_datetime":
			_ = clipboard.WriteAll(carbon.Yesterday().ToDateTimeString())
		case "tomorrow_timestamp":
			_ = clipboard.WriteAll(strconv.FormatInt(carbon.Tomorrow().ToTimestamp(), 10))
		case "tomorrow_date":
			_ = clipboard.WriteAll(carbon.Tomorrow().ToDateString())
		case "tomorrow_datetime":
			_ = clipboard.WriteAll(carbon.Tomorrow().ToDateTimeString())
		}
	})

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

func service(m *gotray.Menu, menu menu) {
	m.SetIcon(getIcon("stop"))
	status := gotray.NewMenu().SetTitle("Not Start")
	start := gotray.NewMenu().SetTitle("Start").SetIcon(getIcon("start"))
	start.Click(func() {
		execCmd(menu.Values[0])
	})
	restart := gotray.NewMenu().SetTitle("Restart").SetIcon(getIcon("Restart"))
	restart.Click(func() {
		execCmd(menu.Values[1])
	})
	stop := gotray.NewMenu().SetTitle("Stop").SetIcon(getIcon("stop"))
	stop.Click(func() {
		execCmd(menu.Values[2])
	})
	var f = func(restart, stop *gotray.Menu) {
		pid := getPid(menu.Pid)
		if pid != "" {
			status.SetTitle("Started(Pid:" + pid + ")")
			m.SetIcon(getIcon("running"))
			//restart.Enable()
			//stop.Enable()
			start.Disable()
		} else {
			m.SetIcon(getIcon("stop"))
			status.SetTitle("Not start")
			//start.Enable()
			stop.Disable()
			restart.Disable()
		}
		status.Disable()
	}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			f(restart, stop)
		}
	}()
}

package main

import (
	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/golang-module/carbon"
	"os/exec"
	"runtime"
	"strconv"
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

func clipboardCopy(item *systray.MenuItem, menu menu) {
	switch menu.Values[0] {
	case "timestamp":
		go func() {
			for {
				select {
				case <-item.ClickedCh:
					_ = clipboard.WriteAll(strconv.FormatInt(carbon.Now().ToTimestamp(), 10))
				}
			}
		}()
	case "date":
		go func() {
			for {
				select {
				case <-item.ClickedCh:
					_ = clipboard.WriteAll(carbon.Now().ToDateString())
				}
			}
		}()
	case "datetime":
		go func() {
			for {
				select {
				case <-item.ClickedCh:
					_ = clipboard.WriteAll(carbon.Now().ToDateTimeString())
				}
			}
		}()
	case "yesterday_timestamp":
		go func() {
			for {
				select {
				case <-item.ClickedCh:
					_ = clipboard.WriteAll(strconv.FormatInt(carbon.Yesterday().ToTimestamp(), 10))
				}
			}
		}()
	case "yesterday_date":
		go func() {
			for {
				select {
				case <-item.ClickedCh:
					_ = clipboard.WriteAll(carbon.Yesterday().ToDateString())
				}
			}
		}()
	case "yesterday_datetime":
		go func() {
			for {
				select {
				case <-item.ClickedCh:
					_ = clipboard.WriteAll(carbon.Yesterday().ToDateTimeString())
				}
			}
		}()
	case "tomorrow_timestamp":
		go func() {
			for {
				select {
				case <-item.ClickedCh:
					_ = clipboard.WriteAll(strconv.FormatInt(carbon.Tomorrow().ToTimestamp(), 10))
				}
			}
		}()
	case "tomorrow_date":
		go func() {
			for {
				select {
				case <-item.ClickedCh:
					_ = clipboard.WriteAll(carbon.Tomorrow().ToDateString())
				}
			}
		}()
	case "tomorrow_datetime":
		go func() {
			for {
				select {
				case <-item.ClickedCh:
					_ = clipboard.WriteAll(carbon.Tomorrow().ToDateTimeString())
				}
			}
		}()
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
	status := item.AddSubMenuItem("Not start", "")
	start := item.AddSubMenuItem("Start", "")
	start.SetIcon(getIcon("start"))
	restart := item.AddSubMenuItem("Restart", "")
	restart.SetIcon(getIcon("restart"))
	restart.Disable()
	stop := item.AddSubMenuItem("Stop", "")
	stop.SetIcon(getIcon("termination"))
	stop.Disable()
	var f = func(restart, stop *systray.MenuItem) {
		pid := getPid(menu.Pid)
		if pid != "" {
			status.SetTitle("Started(Pid:" + pid + ")")
			item.SetIcon(getIcon("running"))
			restart.Enable()
			stop.Enable()
			start.Disable()
		} else {
			item.SetIcon(getIcon("stop"))
			status.SetTitle("Not start")
			start.Enable()
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
	for {
		select {
		case <-item.ClickedCh:
		case <-restart.ClickedCh:
			execCmd(menu.Values[1])
		case <-start.ClickedCh:
			execCmd(menu.Values[0])
		case <-stop.ClickedCh:
			execCmd(menu.Values[2])
		}
	}
}

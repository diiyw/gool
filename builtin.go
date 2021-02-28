package main

import (
	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"strconv"
	"time"
)

func copyTimestamp(item *systray.MenuItem, menu menu) {
	for {
		select {
		case <-item.ClickedCh:
			_ = clipboard.WriteAll(strconv.FormatInt(time.Now().Unix(), 10))
		}
	}
}



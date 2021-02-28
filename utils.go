package main

import "os"

// 根据进程名称获取进程ID
func getPid(pidFile string) string {
	pid, _ := os.ReadFile(pidFile)
	return string(pid)
}


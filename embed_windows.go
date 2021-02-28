package main

import "embed"

//go:embed assets/gool.ico assets/stop.ico assets/running.ico assets/start.ico
var icons embed.FS

const suffix = ".ico"
package main

import "embed"

//go:embed assets/gool.ico assets/stop.ico assets/running.ico assets/start.ico  assets/restart.ico assets/termination.ico
var icons embed.FS

const suffix = ".ico"
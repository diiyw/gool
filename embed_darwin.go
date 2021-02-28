package main

import "embed"

//go:embed assets/gool.png assets/stop.png assets/running.png assets/start.png
var icons embed.FS

const suffix = ".png"
package main

import "embed"

//go:embed assets/*
var icons embed.FS

func getIcon(name string) []byte {
	goolIcon, err := icons.ReadFile("assets/" + name + ".png")
	if err != nil {
		panic(err)
	}
	return goolIcon
}

package main

func getIcon(name string) []byte {
	goolIcon, err := icons.ReadFile("assets/" + name + suffix)
	if err != nil {
		panic(err)
	}
	return goolIcon
}

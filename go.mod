module github.com/diiyw/gool

go 1.16

require (
	github.com/atotto/clipboard v0.1.4
	github.com/diiyw/gotray v0.0.0-20210310160733-d929fbe9779f
	github.com/getlantern/systray v1.1.0
	github.com/golang-module/carbon v1.3.3
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/diiyw/gotray v0.0.0-20210310160733-d929fbe9779f => ../gotray
)
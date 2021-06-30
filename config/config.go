package config

import "os"

var Endpoint = "https://calibration.node.glif.io"
var Token = ""

func init() {
	endpoint := os.Getenv("ENDPOINT")
	if endpoint != "" {
		Endpoint = endpoint
	}
	token := os.Getenv("TOKEN")
	if token != "" {
		Token = token
	}
}

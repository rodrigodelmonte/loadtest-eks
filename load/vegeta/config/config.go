package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Duration int
	Freq     int
	Method   string
	PromAddr string
	URL      string
	Period   time.Duration
	TestName string
}

func NewConfig() *Config {

	duration, _ := strconv.Atoi(os.Getenv("DURATION"))
	freq, _ := strconv.Atoi(os.Getenv("FREQUENCY"))
	return &Config{
		Duration: duration,
		Method:   os.Getenv("METHOD"),
		PromAddr: os.Getenv("PROMETHEUS_ADDR"),
		URL:      os.Getenv("URL"),
		Freq:     freq,
		TestName: os.Getenv("TEST_NAME"),
	}
}

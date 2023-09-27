package main

import (
	"time"
)

const (
	SrvBindAddr  = "0.0.0.0:8080"
	TimeoutWrite = time.Second * 15
	TimeoutRead  = time.Second * 15
	TimeoutIdle  = time.Second * 60
)

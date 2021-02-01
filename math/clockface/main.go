package main

import (
	clockface "TJ-Yusuke/golearning/math"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}

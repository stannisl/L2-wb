package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	currTime, err := ntp.Time("pool.ntp.org")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Current time:", currTime.Format(time.RFC3339))
}

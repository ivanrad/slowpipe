// read standard input, and write to standard output with the specified bitrate
//
// usage: slowpipe [-b bitrate]
//
// exit status: 0 for success, 1 for bitrate out of range, and 2 for
// unrecognized or invalid argument
//

package main

import (
	"flag"
	"io"
	"log"
	"os"
	"time"
)

const (
	minBitrate     = 1
	maxBitrate     = 8 * 1000000000
	defaultBitrate = 300
)

func main() {
	log.SetFlags(0)

	var bitrate int

	flag.IntVar(&bitrate, "b", defaultBitrate, "output bitrate (bit/s)")
	flag.Parse()

	if bitrate < minBitrate || bitrate > maxBitrate {
		log.Fatalf("invalid bitrate: %v (must be between %v and %v)",
			bitrate, minBitrate, maxBitrate)
	}

	sleepDur := time.Duration(maxBitrate/bitrate) * time.Nanosecond
	buf := make([]byte, 1)
	for {
		if _, err := os.Stdin.Read(buf); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if _, err := os.Stdout.Write(buf); err != nil {
			log.Fatal(err)
		}
		time.Sleep(sleepDur)
	}
}

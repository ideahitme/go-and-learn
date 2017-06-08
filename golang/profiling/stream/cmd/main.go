package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"

	challenge "github.com/ideahitme/go-and-learn/golang/profiling/stream"
)

// cpuprofile points to the file where results will be written
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var duration = flag.Int("duration", 1, "how long to profile")

func main() {
	flag.Parse()

	// enable pprof
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile() // this is necessary to flush profiling data to the file
		//memory profiling
		//pprof.Lookup("heap").WriteTo(f)
	}

	//run ../challenge.go Replace func

	var data = []struct {
		input  []byte
		output []byte
	}{
		{[]byte("abc"), []byte("abc")},
		{[]byte("elvis"), []byte("Elvis")},
		{[]byte("aElvis"), []byte("aElvis")},
		{[]byte("abcelvis"), []byte("abcElvis")},
		{[]byte("eelvis"), []byte("eElvis")},
		{[]byte("aelvis"), []byte("aElvis")},
		{[]byte("aabeeeelvis"), []byte("aabeeeElvis")},
		{[]byte("e l v i s"), []byte("e l v i s")},
		{[]byte("aa bb e l v i saa"), []byte("aa bb e l v i saa")},
		{[]byte(" elvi s"), []byte(" elvi s")},
		{[]byte("elvielvis"), []byte("elviElvis")},
		{[]byte("elvielvielviselvi1"), []byte("elvielviElviselvi1")},
		{[]byte("elvielviselvis"), []byte("elviElvisElvis")},
	}

	find := []byte("elvis")
	replace := []byte("Elvis")
	output := &bytes.Buffer{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*duration)*time.Second)
	i := 0
	defer cancel()

	for {
		i++
		i %= len(data)
		select {
		case <-ctx.Done():
			return
		default:
			output.Reset()
			challenge.Replace(data[i].input, find, replace, output)
		}
	}
}

package main

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"testing"
)

var testData []byte

func Benchmark(b *testing.B) {
	testData = []byte{}
	for i := 0; i < 100000; i++ {
		if rand.Intn(2) == 0 {
			testData = append(testData, byte('a'))
		} else {
			testData = append(testData, byte('b'))
		}
	}
	b.Run("With pipe", benchmarkPipe)
	b.Run("With buffer", benchmarkBuffer)
}

func benchmarkPipe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf := bytes.NewReader(testData)
		ioutil.ReadAll(PipedStream(bf))
	}
}

func benchmarkBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf := bytes.NewReader(testData)
		ioutil.ReadAll(BufferedStream(bf))
	}
}

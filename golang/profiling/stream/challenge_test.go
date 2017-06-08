package challenge

import (
	"bytes"
	"testing"
)

func TestReplace(t *testing.T) {
	find := []byte("yerken")
	replace := []byte("Yerken")
	for _, ti := range []struct {
		title  string
		input  []byte
		output []byte
	}{
		{
			title:  "simple",
			input:  []byte("yerken"),
			output: []byte("Yerken"),
		},
		{
			title:  "more sophisticated",
			input:  []byte("yerken Yerken dyerken"),
			output: []byte("Yerken Yerken dYerken"),
		},
		{
			title:  "without changes",
			input:  []byte("y e r k e n"),
			output: []byte("y e r k e n"),
		},
		{
			title:  "overlapping",
			input:  []byte("yeryerken"),
			output: []byte("yerYerken"),
		},
		{
			title:  "overlapping no.2",
			input:  []byte("yerkenyer"),
			output: []byte("Yerkenyer"),
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			output := bytes.NewBuffer([]byte{})
			Replace(ti.input, find, replace, output)
			if bytes.Compare(ti.output, output.Bytes()) != 0 {
				t.Errorf("expected: %v, got: %v", ti.output, output.Bytes())
			}
		})
	}
}

func TestReplaceProblematic(t *testing.T) {
	find := []byte("yerken")
	replace := []byte("Yerken")
	for _, ti := range []struct {
		title  string
		input  []byte
		output []byte
	}{
		{
			title:  "simple",
			input:  []byte("yerken"),
			output: []byte("Yerken"),
		},
		{
			title:  "more sophisticated",
			input:  []byte("yerken Yerken dyerken"),
			output: []byte("Yerken Yerken dYerken"),
		},
		{
			title:  "without changes",
			input:  []byte("y e r k e n"),
			output: []byte("y e r k e n"),
		},
		{
			title:  "overlapping",
			input:  []byte("yeryerken"),
			output: []byte("yerYerken"),
		},
		{
			title:  "overlapping no.2",
			input:  []byte("yerkenyer"),
			output: []byte("Yerkenyer"),
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			output := bytes.NewBuffer([]byte{})
			ReplaceProblematic(ti.input, find, replace, output)
			if bytes.Compare(ti.output, output.Bytes()) != 0 {
				t.Errorf("expected: %v, got: %v", ti.output, output.Bytes())
			}
		})
	}
}
func TestReplaceNonStream(t *testing.T) {
	find := []byte("yerken")
	replace := []byte("Yerken")
	for _, ti := range []struct {
		title  string
		input  []byte
		output []byte
	}{
		{
			title:  "simple",
			input:  []byte("yerken"),
			output: []byte("Yerken"),
		},
		{
			title:  "more sophisticated",
			input:  []byte("yerken Yerken dyerken"),
			output: []byte("Yerken Yerken dYerken"),
		},
		{
			title:  "without changes",
			input:  []byte("y e r k e n"),
			output: []byte("y e r k e n"),
		},
		{
			title:  "overlapping",
			input:  []byte("yeryerken"),
			output: []byte("yerYerken"),
		},
		{
			title:  "overlapping no.2",
			input:  []byte("yerkenyer"),
			output: []byte("Yerkenyer"),
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			output := bytes.NewBuffer([]byte{})
			ReplaceNonStream(ti.input, find, replace, output)
			if bytes.Compare(ti.output, output.Bytes()) != 0 {
				t.Errorf("expected: %v, got: %v", ti.output, output.Bytes())
			}
		})
	}
}

func TestReplaceTricky(t *testing.T) {
	find := []byte("abaa")
	replace := []byte("ABAA")
	data := []byte("ababaa")
	output := bytes.NewBuffer([]byte{})
	Replace(data, find, replace, output)
	if bytes.Compare(output.Bytes(), []byte("abABAA")) != 0 {
		t.Fatalf("incorrect result: %v", output.Bytes())
	}
}

func TestReplaceProblematicTricky(t *testing.T) {
	find := []byte("abaa")
	replace := []byte("ABAA")
	data := []byte("ababaa")
	output := bytes.NewBuffer([]byte{})
	ReplaceProblematic(data, find, replace, output)
	if bytes.Compare(output.Bytes(), []byte("abABAA")) != 0 {
		t.Fatalf("incorrect result: %v", output.Bytes())
	}
}

func TestReplaceNonStreamTricky(t *testing.T) {
	find := []byte("abaa")
	replace := []byte("ABAA")
	data := []byte("ababaa")
	output := bytes.NewBuffer([]byte{})
	ReplaceNonStream(data, find, replace, output)
	if bytes.Compare(output.Bytes(), []byte("abABAA")) != 0 {
		t.Fatalf("incorrect result: %v", output.Bytes())
	}
}

func BenchmarkReplace(b *testing.B) {
	find := []byte("elvis")
	replace := []byte("Elvis")
	data := assembleInputStream()
	output := bytes.Buffer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output.Reset()
		Replace(data, find, replace, &output)
	}
}

func BenchmarkReplaceProblematic(b *testing.B) {
	find := []byte("elvis")
	replace := []byte("Elvis")
	data := assembleInputStream()
	output := bytes.Buffer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output.Reset()
		ReplaceProblematic(data, find, replace, &output)
	}
}

func BenchmarkReplaceNonStream(b *testing.B) {
	find := []byte("elvis")
	replace := []byte("Elvis")
	data := assembleInputStream()
	output := bytes.Buffer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output.Reset()
		ReplaceNonStream(data, find, replace, &output)
	}
}

func assembleInputStream() []byte {
	var in []byte
	for _, d := range data {
		in = append(in, d.input...)
	}
	return in
}

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

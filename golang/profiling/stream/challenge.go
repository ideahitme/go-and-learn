package challenge

import (
	"bytes"
	"io"
)

// Replace given data it replaces all occurences of find with replace
func Replace(data, find, replace []byte, output *bytes.Buffer) {
	stream := bytes.NewReader(data)
	size := len(find)
	i := 0

	output.Grow(int(stream.Size()))

	var b byte
	var err error

	for {
		b, err = stream.ReadByte()
		if err != nil {
			break
		}
		if find[i] == b {
			if i == size-1 {
				output.Write(replace)
				i = 0
			} else {
				i++
			}
			continue
		}

		stream.Seek(-int64(i+1), io.SeekCurrent)
		b, _ = stream.ReadByte()
		output.WriteByte(b)

		i = 0
	}
	//write whatever is left
	output.Write(find[:i])
}

// ReplaceProblematic worse version of the Replace
func ReplaceProblematic(data, find, replace []byte, output *bytes.Buffer) {
	stream := bytes.NewReader(data)
	size := len(find)

	var b byte
	var err error
	var n int
	buf := make([]byte, size, size)

	for {
		n, err = io.ReadFull(stream, buf)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				stream.Seek(int64(-n), io.SeekCurrent)
			}
			break
		}

		if bytes.Compare(buf, find) == 0 {
			output.Write(replace)
			continue
		}

		stream.Seek(-int64(size), io.SeekCurrent)
		b, _ = stream.ReadByte()
		output.WriteByte(b)
	}
	//write whatever is left
	stream.WriteTo(output)
}

// ReplaceNonStream non stream version
func ReplaceNonStream(data, find, replace []byte, output *bytes.Buffer) {
	size := len(find)
	total := len(data)
	i := 0
	j := 0
	var b byte

	for j < total {
		b = data[j]

		if find[i] == b {
			if i == size-1 {
				output.Write(replace)
				i = 0
			} else {
				i++
			}
			j++
			continue
		}

		// roll back
		j -= i
		output.WriteByte(data[j])

		j++
		i = 0
	}

	output.Write(find[:i])
}

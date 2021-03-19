package encode_and_decode_strings_271

import (
	"bytes"
	"encoding/binary"
)

type Codec struct {
}

func (codec *Codec) Encode(strs []string) string {
	var offsets [][]int

	var buffer bytes.Buffer
	for _, str := range strs {
		begin := buffer.Len()
		buffer.WriteString(str)
		offsets = append(offsets, []int{begin, buffer.Len()})
	}

	var head bytes.Buffer

	var offsetBuffer = make([]byte, 4)
	binary.BigEndian.PutUint32(offsetBuffer, uint32(len(offsets)))

	head.Write(offsetBuffer)

	for _, offset := range offsets {
		var offsetBuffer = make([]byte, 4)
		binary.BigEndian.PutUint32(offsetBuffer, uint32(offset[0]))
		head.Write(offsetBuffer)

		binary.BigEndian.PutUint32(offsetBuffer, uint32(offset[1]))
		head.Write(offsetBuffer)
	}

	head.Write(buffer.Bytes())

	return head.String()
}

func (codec Codec) Decode(str string) []string {
	data := []byte(str)
	offsetSize := binary.BigEndian.Uint32(data)
	data = data[4:]

	var offsets [][]int
	for i := 0; i < int(offsetSize); i++ {
		begin := binary.BigEndian.Uint32(data)
		data = data[4:]

		end := binary.BigEndian.Uint32(data)
		data = data[4:]

		offsets = append(offsets, []int{int(begin), int(end)})
	}
	var results []string
	for _, offset := range offsets {
		results = append(results, string(data[offset[0]:offset[1]]))
	}

	return results
}

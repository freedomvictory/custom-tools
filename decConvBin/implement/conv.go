package implement

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func comTypeToBytes(i interface{}) []byte {

	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, i)
	if err != nil {
		log.Printf("[comTypeToBytes] (err) Fail to convert:%v", err)
		return nil
	}
	return buf.Bytes()

}

func InsertNth(s string, n int, in string) string {

	var n_1 = n - 1
	var l_1 = len(s) - 1
	var buffer bytes.Buffer
	for i, rune := range s {

		buffer.WriteRune(rune)

		//out[j] = s[i]
		if i%n == n_1 && i != l_1 {
			for _, rune := range in {
				buffer.WriteRune(rune)
			}
		}
	}
	return buffer.String()
}

func MuxDecToBinHandler(data interface{}, typ string, format int) (out string, err error) {

	bytes := []byte{}
	switch typ {

	case "f64":
		data_f64 := data.(float64)
		bytes = comTypeToBytes(data_f64)
	case "f32":
		data_f32 := data.(float32)
		bytes = comTypeToBytes(data_f32)
	case "i32":
		data_i32 := data.(int32)
		bytes = comTypeToBytes(data_i32)
	case "i16":
		data_i16 := data.(int16)
		bytes = comTypeToBytes(data_i16)
	}
	if bytes == nil {
		return "", fmt.Errorf("convert to binay fail")
	}

	bytes_s := fmt.Sprintf("%x", bytes)
	switch format {
	case 1:
		out = fmt.Sprintf("%#x", bytes)

	case 2:
		out = InsertNth(bytes_s, 2, ",0x")
		out = "0x" + out

	case 3:
		out = InsertNth(bytes_s, 2, "\\x")
		out = "\\x" + out

	default:

	}

	return out, nil

}

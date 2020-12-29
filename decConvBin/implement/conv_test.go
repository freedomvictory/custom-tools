package implement

import (
	"testing"
)

type inputFormatData struct {
	data      interface{}
	typ       string
	outFormat int
}

func TestInsetNth(t *testing.T) {
	out := InsertNth("helloworldhelloworldhelloworld", 5, "-")
	out_2 := InsertNth("2345761a2b", 2, ",")
	out_3 := InsertNth("2345761a2b", 2, "/x")

	t.Logf("out :%v\n", out)
	t.Logf("out :%v\n", out_2)
	t.Logf("out :%v\n", out_3)

	t.FailNow()
}
func TestMuxDecToBin(t *testing.T) {

	iFDs := []inputFormatData{
		{
			data:      int32(25),
			typ:       "i32",
			outFormat: 1,
		},
		{
			data:      float32(2.55),
			typ:       "f32",
			outFormat: 2,
		},
		{
			data:      float64(7.88),
			typ:       "f64",
			outFormat: 3,
		},
	}

	for _, val := range iFDs {
		out, err := MuxDecToBinHandler(val.data, val.typ, val.outFormat)
		if err != nil {
			t.Fatalf("Fail to call MuxDecToBinHandler : %v", err)
		}
		t.Logf("[TestMuxDecToBin] (log) out:%v", out)
	}
	t.FailNow()

}

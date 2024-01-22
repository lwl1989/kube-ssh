package timex

import "testing"

func TestDateStrTime_Scan(t *testing.T) {

	d := &DateStrTime{}
	d.Scan("2022-01-02")
	t.Log(d)
	d.Scan("2022-01-02 10:11:12")
	t.Log(d)
	d.Scan([]byte("2022-01-02 10:11:12"))
	t.Log(d)
	d.Scan([]byte("2022-01-02"))
	t.Log(d)
	d.Scan(1641089472)
	t.Log(d)
	d.Scan("1641089472")
	t.Log(d)
	d.Scan([]byte("1641089472"))
	t.Log(d)
}

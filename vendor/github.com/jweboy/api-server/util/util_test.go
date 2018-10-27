package util

import "testing"

func TestGenShortId(t *testing.T) {
	shortID, err := GenShortId()

	if shortID == "" || err != nil {
		t.Error("GenShortId failed.")
	}

	t.Log("GenShortId test 	pass.")
}

func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	// 停止测试时间
	b.StopTimer()

	shortID, err := GenShortId()

	if shortID == "" || err != nil {
		b.Error(err)
	}

	// 重新开始时间
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

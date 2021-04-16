package kgo

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNumber_NumberFormat(t *testing.T) {
	var res string

	res = KNum.NumberFormat(floNum1, 10, ".", "")
	assert.Equal(t, "12345.1234567890", res)

	res = KNum.NumberFormat(floNum2, 6, ".", ",")
	assert.Equal(t, "12,345,678.123457", res)

	res = KNum.NumberFormat(floNum3, 0, ".", "")
	assert.Equal(t, "-123", res)

	res = KNum.NumberFormat(math.Pi, 15, ".", "")
	assert.Equal(t, "3.141592653589793", res)
}

func BenchmarkNumber_NumberFormat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.NumberFormat(floNum1, 3, ".", "")
	}
}

func TestNumber_Range(t *testing.T) {
	var res []int
	var start, end int

	//升序
	start, end = 1, 5
	res = KNum.Range(start, end)
	assert.Equal(t, 5, len(res))
	assert.Equal(t, start, res[0])

	//降序
	start, end = 5, 1
	res = KNum.Range(start, end)
	assert.Equal(t, 5, len(res))
	assert.Equal(t, start, res[0])

	//起始和结尾相同
	start, end = 3, 3
	res = KNum.Range(start, end)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, start, res[0])
}

func BenchmarkNumber_Range(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Range(0, 9)
	}
}

func TestNumber_AbsFloat(t *testing.T) {
	var res float64

	res = KNum.AbsFloat(floNum3)
	assert.Greater(t, res, 0.0)
}

func BenchmarkNumber_AbsFloat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.AbsFloat(floNum3)
	}
}

func TestNumber_AbsInt(t *testing.T) {
	var res int64

	res = KNum.AbsInt(-123)
	assert.Greater(t, res, int64(0))
}

func BenchmarkNumber_AbsInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.AbsInt(-123)
	}
}

func TestNumber_FloatEqual(t *testing.T) {
	var res bool

	//默认小数位
	res = KNum.FloatEqual(floNum1, floNum4)
	assert.True(t, res)

	res = KNum.FloatEqual(floNum1, floNum4, 0)
	assert.True(t, res)

	res = KNum.FloatEqual(floNum1, floNum4, 11)
	assert.True(t, res)

	res = KNum.FloatEqual(floNum1, floNum4, 12)
	assert.False(t, res)
}

func BenchmarkNumber_FloatEqual(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.FloatEqual(floNum1, floNum4)
	}
}

func TestNumber_RandInt64(t *testing.T) {
	var min, max, res int64

	min, max = -9, 9
	res = KNum.RandInt64(min, max)
	assert.GreaterOrEqual(t, res, min)
	assert.LessOrEqual(t, res, max)

	res = KNum.RandInt64(max, max)
	assert.Equal(t, res, max)
}

func TestNumber_RandInt64_PanicMin(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var min, max int64
	min, max = 9, -9
	KNum.RandInt64(min, max)
}

func TestNumber_RandInt64_PanicOverflow(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	KNum.RandInt64(INT64_MIN, INT64_MAX)
}

func BenchmarkNumber_RandInt64(b *testing.B) {
	b.ResetTimer()
	var min, max int64
	min, max = 9, -9
	for i := 0; i < b.N; i++ {
		KNum.RandInt64(min, max)
	}
}

func TestNumber_RandInt(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res int
	min, max := -9, 9
	res = KNum.RandInt(min, max)
	assert.GreaterOrEqual(t, res, min)
	assert.LessOrEqual(t, res, max)

	res = KNum.RandInt(max, max)
	assert.Equal(t, res, max)

	KNum.RandInt(9, 1)
}

func BenchmarkNumber_RandInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.RandInt(-9, 9)
	}
}

func TestNumber_Rand(t *testing.T) {
	var res int
	min, max := -9, 9
	res = KNum.Rand(min, max)
	assert.GreaterOrEqual(t, res, min)
	assert.LessOrEqual(t, res, max)

	res = KNum.Rand(max, max)
	assert.Equal(t, res, max)
}

func BenchmarkNumber_Rand(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Rand(-9, 9)
	}
}

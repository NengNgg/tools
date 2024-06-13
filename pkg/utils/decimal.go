package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func Float32ToFloat64(f float32) float64 {
	r, _ := decimal.NewFromFloat32(f).Float64()
	return r
}
func IntToFloat64(i int) (res float64) {
	res, _ = decimal.NewFromInt(int64(i)).Float64()
	return
}
func MulFlat64(f1, f2 float64) (res float64) {
	f1D := decimal.NewFromFloat(f1)
	f2D := decimal.NewFromFloat(f2)
	res, _ = f1D.Mul(f2D).Float64()
	return
}

// DivFlat64 f1/f2
func DivFlat64(f1, f2 float64) (res float64) {
	if f2 == 0 {
		return 0
	}
	f1D := decimal.NewFromFloat(f1)
	f2D := decimal.NewFromFloat(f2)
	res, _ = f1D.Div(f2D).Float64()
	return
}
func AddFlat64(f1, f2 float64) (res float64) {
	f1D := decimal.NewFromFloat(f1)
	f2D := decimal.NewFromFloat(f2)
	res, _ = f1D.Add(f2D).Float64()
	return
}

// SubFlat64 f1-f2
func SubFlat64(f1, f2 float64) (res float64) {
	f1D := decimal.NewFromFloat(f1)
	f2D := decimal.NewFromFloat(f2)
	res, _ = f1D.Sub(f2D).Float64()
	return
}

func SubString(fs1, fs2 string) (res string, err error) {
	f1, err := decimal.NewFromString(fs1)
	if err != nil {
		return "", err
	}
	f2, err := decimal.NewFromString(fs2)
	if err != nil {
		return "", err
	}

	res = f1.Sub(f2).String()
	return
}
func MulString(fs1, fs2 string) (res string, err error) {
	f1, err := decimal.NewFromString(fs1)
	if err != nil {
		return "", err
	}
	f2, err := decimal.NewFromString(fs2)
	if err != nil {
		return "", err
	}

	res = f1.Mul(f2).String()
	return
}

func AddString(fs1, fs2 string) (res string, err error) {
	f1, err := decimal.NewFromString(fs1)
	if err != nil {
		return "", err
	}
	f2, err := decimal.NewFromString(fs2)
	if err != nil {
		return "", err
	}

	res = f1.Add(f2).String()
	return
}

// DivString b保留位数
func DivString(fs1, fs2 string, b int32) (res string, err error) {
	f1, err := decimal.NewFromString(fs1)
	if err != nil {
		return "", err
	}
	f2, err := decimal.NewFromString(fs2)
	if err != nil {
		return "", err
	}

	if f2.IsZero() {
		return "", fmt.Errorf("division by zero")
	}

	// 保留两位小数
	res = f1.Div(f2).Round(b).String()
	return
}

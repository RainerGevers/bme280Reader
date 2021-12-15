package readings

import (
	"github.com/RainerGevers/bme280Reader/calibration"
	"math"
)

const (
	P1Addr = 0xF7
	P2Addr = 0xF8
	P3Addr = 0xF9

	T1Addr = 0xFA
	T2Addr = 0xFB
	T3Addr = 0xFC

	H1Addr = 0xFD
	H2Addr = 0xFE
)

type UncompensatedReading struct {
	Pressure    uint32
	Temperature uint32
	Humidity    uint32
}

type CompensatedReading struct {
	Pressure    float64
	Temperature float64
	Humidity    float64
}

func Read(methodCall func(uint8) uint32) UncompensatedReading {
	uncompRead := UncompensatedReading{}

	uncompRead.Temperature = (methodCall(T1Addr)<<16 | methodCall(T2Addr)<<8 | methodCall(T3Addr)) >> 4
	uncompRead.Pressure = (methodCall(P1Addr)<<16 | methodCall(P2Addr)<<8 | methodCall(P3Addr)) >> 4
	uncompRead.Humidity = methodCall(H1Addr)<<8 | methodCall(H2Addr)

	return uncompRead
}

func CalcTFine(temp uint32) float64 {
	v1 := (float64(temp)/16384.0 - float64(calibration.CalParams.DigT1)/1024.0) * float64(calibration.CalParams.DigT2)
	v2 := math.Pow(float64(float64(temp)/131072.0-float64(calibration.CalParams.DigT1)/8192.0), 2.0) * float64(calibration.CalParams.DigT3)
	return v1 + v2
}

func CalcHumidity(humidity uint32, tFine float64) float64 {
	res := tFine - 76800.0
	res = (float64(humidity) - (float64(calibration.CalParams.DigH4)*64.0 + float64(calibration.CalParams.DigH5)/16384.0*res)) * (float64(calibration.CalParams.DigH2) / 65536.0 * (1.0 + float64(calibration.CalParams.DigH6)/67108864.0*res*(1.0+float64(calibration.CalParams.DigH3)/67108864.0*res)))
	res = res * (1.0 - (float64(calibration.CalParams.DigH1) * res / 524288.0))
	return max(0.0, min(res, 100.0))
}

func CalcPressure(pressure uint32, tFine float64) float64 {
	v1 := (tFine / 2.0) - 64000.0
	v2 := v1 * v1 * (float64(calibration.CalParams.DigP6) / 32768.0)
	v2 = v2 + v1 + (float64(calibration.CalParams.DigP5) * 2.0)
	v2 = (v2 / 4.0) + (float64(calibration.CalParams.DigP4) * 65536.0)
	v1 = ((float64(calibration.CalParams.DigP3)*v1*v1)/524288.0 + (float64(calibration.CalParams.DigP2) * v1)) / 524288.0
	v1 = (1.0 + v1/32768.0) * float64(calibration.CalParams.DigP1)

	if v1 == 0.0 {
		return 0.0
	}

	res := 1048576.0 - float64(pressure)
	res = ((res - v2/4096.0) * 6250.0) / v1
	v1 = (float64(calibration.CalParams.DigP9) * res * res) / 2147483648.0
	v2 = res * (float64(calibration.CalParams.DigP8) / 32768.0)
	res = res + (v1+v2+float64(calibration.CalParams.DigP7))/16.0

	return res
}

func CompensateReading(uncompedReading UncompensatedReading) CompensatedReading {
	comp := CompensatedReading{}
	tFine := CalcTFine(uncompedReading.Temperature)
	comp.Temperature = tFine / 5120.0
	comp.Humidity = CalcHumidity(uncompedReading.Humidity, tFine)
	comp.Pressure = CalcPressure(uncompedReading.Pressure, tFine)

	return comp
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

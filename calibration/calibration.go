package calibration

var CalParams Params

const (
	DigT1Addr = 0x88
	DigT2Addr = 0x8A
	DigT3Addr = 0x8C

	DigP1Addr = 0x8E
	DigP2Addr = 0x90
	DigP3Addr = 0x92
	DigP4Addr = 0x94
	DigP5Addr = 0x96
	DigP6Addr = 0x98
	DigP7Addr = 0x9A
	DigP8Addr = 0x9C
	DigP9Addr = 0x9E

	DigH1Addr = 0xA1
	DigH2Addr = 0xE1
	DigH3Addr = 0xE3
	DigH6Addr = 0xE7

	E4Addr = 0xE4
	E5Addr = 0xE5
	E6Addr = 0xE6
)

type Params struct {
	DigT1        uint16
	DigT2, DigT3 int16

	DigP1                                                  uint16
	DigP2, DigP3, DigP4, DigP5, DigP6, DigP7, DigP8, DigP9 int16

	DigH1 uint8
	DigH2 int16
	DigH3 int8
	DigH4 int8
	DigH5 int8
	DigH6 int8

	E4, E5, E6 int8
}

func SetCalParams(methodCall func(uint16) uint16) {
	CalParams = Params{}
	CalParams.DigT1 = methodCall(DigT1Addr)
	CalParams.DigT2 = int16(methodCall(DigT2Addr))
	CalParams.DigT3 = int16(methodCall(DigT3Addr))

	CalParams.DigP1 = methodCall(DigP1Addr)
	CalParams.DigP2 = int16(methodCall(DigP2Addr))
	CalParams.DigP3 = int16(methodCall(DigP3Addr))
	CalParams.DigP4 = int16(methodCall(DigP4Addr))
	CalParams.DigP5 = int16(methodCall(DigP5Addr))
	CalParams.DigP6 = int16(methodCall(DigP6Addr))
	CalParams.DigP7 = int16(methodCall(DigP7Addr))
	CalParams.DigP8 = int16(methodCall(DigP8Addr))
	CalParams.DigP9 = int16(methodCall(DigP9Addr))

	CalParams.DigH1 = uint8(methodCall(DigH1Addr))
	CalParams.DigH2 = int16(methodCall(DigH2Addr))
	CalParams.DigH3 = int8(methodCall(DigH3Addr))

	CalParams.E4 = int8(methodCall(E4Addr))
	CalParams.E5 = int8(methodCall(E5Addr))
	CalParams.E6 = int8(methodCall(E6Addr))

	CalParams.DigH4 = CalParams.E4<<4 | CalParams.E5&0x0F
	CalParams.DigH5 = ((CalParams.E5 >> 4) & 0x0F) | (CalParams.E6 << 4)
	CalParams.DigH6 = int8(methodCall(DigH6Addr))
}

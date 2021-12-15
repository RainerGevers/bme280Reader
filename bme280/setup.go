package bme280

const (
	ctrlAddr    = 0xF4
	ctrlHumAddr = 0xF2
)

func SetConfig(methodCall func(uint8, uint8) bool, ctrl, ctrlHum uint8) {
	methodCall(ctrlAddr, ctrl)
	methodCall(ctrlHumAddr, ctrlHum)
}

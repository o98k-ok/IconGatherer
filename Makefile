all:


battery: tool/merge/main.go
	go run tool/merge/main.go ./icon/raw/AirPods ./icon/tool/batteryRing ./icon/batteryLevel/AirPods
	go run tool/merge/main.go ./icon/raw/AppleWatch ./icon/tool/batteryRing ./icon/batteryLevel/AppleWatch
	go run tool/merge/main.go ./icon/raw/iPad ./icon/tool/batteryRing ./icon/batteryLevel/iPad
	go run tool/merge/main.go ./icon/raw/iPhone ./icon/tool/batteryRing ./icon/batteryLevel/iPhone
	go run tool/merge/main.go ./icon/raw/Mac ./icon/tool/batteryRing ./icon/batteryLevel/Mac
	go run tool/merge/main.go ./icon/raw/MagicKeyboard ./icon/tool/batteryRing ./icon/batteryLevel/MagicKeyboard
	go run tool/merge/main.go ./icon/raw/MagicTrackpad ./icon/tool/batteryRing ./icon/batteryLevel/MagicTrackpad
	go run tool/merge/main.go ./icon/raw/mouse ./icon/tool/batteryRing ./icon/batteryLevel/mouse
	go run tool/merge/main.go ./icon/raw/bluetooth ./icon/tool/batteryRing ./icon/batteryLevel/bluetooth
	go run tool/merge/main.go ./icon/raw/VisionPro ./icon/tool/batteryRing ./icon/batteryLevel/VisionPro

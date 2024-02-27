wasm:
	tinygo build --target=wasm -o code/wasm/html/wasm.wasm main.go

server-wasm:
	 go run code/wasm/server.go

blinky:
	tinygo flash --target=arduino-nano33 code/blinky/main.go

blinky2:
	tinygo flash --target=arduino-nano33 code/blinky2/main.go

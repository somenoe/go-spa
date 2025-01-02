wasm:
	set GOARCH=wasm
	set GOOS=js
	go build -o dist/web/app.wasm

build:
	go build

run: build wasm
	./demo.exe

build-wsl:
	GOARCH=wasm GOOS=js go build -o dist/web/app.wasm
	go build

run-wsl: build-wsl
	./demo

tw:
	npx tailwindcss -i input.css -o dist/output.css --watch

serve:
	npx serve dist

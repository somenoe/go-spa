wasm:
	GOARCH=wasm GOOS=js go build -o dist/web/app.wasm

build: wasm
	go build -o tmp/bin/demo

tw:
	npx tailwindcss -i input.css -o dist/output.css

run: build
	./demo

serve:
	npx serve dist

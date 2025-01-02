build:
	GOARCH=wasm GOOS=js go build -o dist/web/app.wasm
	go build

run: build
	./demo

tw:
	npx tailwindcss -i input.css -o dist/output.css --watch

serve:
	npx serve dist

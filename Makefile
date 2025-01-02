build:
	GOARCH=wasm GOOS=js go build -o dist/web/app.wasm
	go build & npx tailwindcss -i input.css -o dist/output.css

tw:
	npx tailwindcss -i input.css -o dist/output.css

run: build
	./demo

serve:
	npx serve dist

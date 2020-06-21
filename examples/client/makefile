build:
	GOARCH=wasm GOOS=js go build -ldflags '-X "main.applicationID=${applicationID}" -X "main.apiToken=${apiToken}"' -o app.wasm

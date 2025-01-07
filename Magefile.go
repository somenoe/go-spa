//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Wasm() error {
	return sh.RunWith(map[string]string{
		"GOARCH": "wasm",
		"GOOS":   "js",
	}, "go", "build", "-o", "dist/web/app.wasm")
}

func Build() error {
	mg.Deps(Wasm)
	return sh.Run("go", "build", "-o", "tmp/bin/main.exe")
	// return sh.Run("go", "run",".")
}

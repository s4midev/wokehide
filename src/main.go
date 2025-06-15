package main

import (
	"fmt"
	"os/exec"

	"github.com/alexflint/go-arg"
	"github.com/wabarc/go-catbox"
)

type args struct {
	Positional string `arg:"positional" help:"The file path or the base64"`
}

func main() {
	var a args

	arg.MustParse(&a)

	fmt.Println(a.Positional)

	if IsBase64(a.Positional) {
		exec.Command("xdg-open", DecodeBase64(a.Positional)).Start()
	} else {
		if url, err := catbox.New(nil).Upload(a.Positional); err != nil {
			fmt.Println(err)
		} else {
			exec.Command("wl-copy", EncodeBase64(url)).Start()
		}
	}
}

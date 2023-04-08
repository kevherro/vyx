// MIT License
//
// Copyright (c) 2023 Kevin Herro
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

package driver

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kevherro/vyx/internal/plugin"
)

func setDefaults(o *plugin.Options) *plugin.Options {
	d := &plugin.Options{}
	if o != nil {
		*d = *o
	}
	if d.Writer == nil {
		d.Writer = writer{}
	}
	if d.UI == nil {
		d.UI = &stdUI{r: bufio.NewReader(os.Stdin)}
	}
	return d
}

type stdUI struct {
	r *bufio.Reader
}

func (ui *stdUI) ReadLine(prompt string) (string, error) {
	os.Stdout.WriteString(prompt)
	return ui.r.ReadString('\n')
}

func (ui *stdUI) Print(args ...any) {
	ui.fPrintf(os.Stderr, args)
}

func (ui *stdUI) PrintErr(args ...any) {
	ui.fPrintf(os.Stderr, args)
}

func (ui *stdUI) fPrintf(f *os.File, args []any) {
	text := fmt.Sprint(args...)
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	f.WriteString(text)
}

// writer implements the Writer interface using a regular file.
type writer struct{}

func (writer) Open(name string) (io.WriteCloser, error) {
	f, err := os.Create(name)
	return f, err
}

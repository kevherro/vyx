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

package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/chzyer/readline"
	"github.com/kevherro/vyx/driver"
)

func main() {
	if err := driver.Vyx(&driver.Options{UI: newUI()}); err != nil {
		fmt.Fprintf(os.Stderr, "vyx: %v\n", err)
		os.Exit(2)
	}
}

type readlineUI struct {
	rl *readline.Instance
}

func newUI() driver.UI {
	rl, err := readline.New("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "readline: %v", err)
		return nil
	}
	return &readlineUI{
		rl: rl,
	}
}

// ReadLine returns a line of text (a command) read from the user.
// prompt is printed before reading the command.
func (r *readlineUI) ReadLine(prompt string) (string, error) {
	r.rl.SetPrompt(prompt)
	return r.rl.Readline()
}

// Print shows a message to the user.
// It is printed over stderr as stdout is reserved for regular output.
func (r *readlineUI) Print(args ...any) {
	text := fmt.Sprint(args...)
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	fmt.Fprint(r.rl.Stderr(), text)
}

// PrintErr shows a message to the user, colored in red for emphasis.
// It is printed over stderr as stdout is reserved for regular output.
func (r *readlineUI) PrintErr(args ...any) {
	text := fmt.Sprint(args...)
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	if readline.IsTerminal(syscall.Stderr) {
		text = colorize(text)
	}
	fmt.Fprint(r.rl.Stderr(), text)
}

// colorize the msg using ANSI color escapes.
func colorize(msg string) string {
	var red = 31
	var colorEscape = fmt.Sprintf("\033[0;%dm", red)
	var colorResetEscape = "\033[0m"
	return colorEscape + msg + colorResetEscape
}

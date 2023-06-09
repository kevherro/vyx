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
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/kevherro/vyx/internal/plugin"
)

var commentStart = "//:" // Sentinel for comments on options.

func interactive(o *plugin.Options) error {
	// Enter the command processing loop.
	greetings(o.UI)
	for {
		input, err := o.UI.ReadLine("(vyx) ")
		if err != nil {
			if err != io.EOF {
				return err
			}
			if input == "" {
				return nil
			}
		}

		// Process assignments of the form variable=value.
		if s := strings.SplitN(input, "=", 2); len(s) > 0 {
			name := strings.TrimSpace(s[0])
			var value string
			if len(s) == 2 {
				value = s[1]
				if comment := strings.LastIndex(value, commentStart); comment != -1 {
					value = value[:comment]
				}
				value = strings.TrimSpace(value)
			}
			if isConfigurable(name) {
				// All non-bool options require inputs.
				if len(s) == 1 && !isBoolConfig(name) {
					o.UI.PrintErr(fmt.Errorf("please specify a value, e.g. %s=<val>", name))
					continue
				}
				if err := configure(name, value); err != nil {
					o.UI.PrintErr(err)
				}
				continue
			}
		}

		tokens := strings.Fields(input)
		if len(tokens) == 0 {
			continue
		}

		switch tokens[0] {
		case "o", "options":
			printCurrentOptions(o.UI)
			continue
		case "help":
			commandHelp(strings.Join(tokens[1:], " "), o.UI)
			continue
		case "exit", "quit", "q":
			return nil
		}

		reply, err := parseTokens(tokens)
		if err == nil {
			o.UI.Print(strings.Join(reply, " "))
		}
		if err != nil {
			o.UI.PrintErr(err)
		}
	}
}

func greetings(ui plugin.UI) {
	ui.Print(`Entering interactive mode (type "help" for commands, "o" for options)`)
}

func printCurrentOptions(ui plugin.UI) {
	var args []string
	c := currentConfig()
	for _, f := range configFields {
		n := f.name
		v := c.get(f)
		comment := ""
		switch {
		case len(f.choices) > 0:
			values := append([]string{}, f.choices...)
			sort.Strings(values)
			comment = "[" + strings.Join(values, " | ") + "]"
		case n == "temperature" && v == "1":
			comment = "default"
		case n == "max_tokens" && v == "2147483647":
			comment = "default"
		case v == "":
			// Add quotes for empty values.
			v = `""`
		}
		if comment != "" {
			comment = commentStart + " " + comment
		}
		args = append(args, fmt.Sprintf("  %-25s = %-20s %s", n, v, comment))
	}
	sort.Strings(args)
	ui.Print(strings.Join(args, "\n"))
}

func commandHelp(args string, ui plugin.UI) {
	ui.Print(args)
}

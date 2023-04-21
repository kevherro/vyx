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
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// config holds settings for a single named config.
// The JSON tag name for a field is used both for JSON
// encoding and as a named variable.
type config struct {
	// Filename for file-based output formats, stdout by default.
	Output string `json:"-"`

	// OpenAI API options.
	Endpoint    string  `json:"endpoint,omitempty"`    // The OpenAI endpoint to use.
	MaxTokens   int     `json:"max_tokens,omitempty"`  // The maximum number of tokens to generate in the completion.
	Temperature float64 `json:"temperature,omitempty"` // What sampling temperature to use, between 0 and 2.
}

// fieldPtr returns a pointer to the field identified by f in c.
func (c *config) fieldPtr(f configField) any {
	return reflect.ValueOf(c).Elem().FieldByIndex(f.field.Index).Addr().Interface()
}

// get returns the value of field f in c.
func (c *config) get(f configField) string {
	switch ptr := c.fieldPtr(f).(type) {
	case *string:
		return *ptr
	case *bool:
		return fmt.Sprint(*ptr)
	case *int:
		return fmt.Sprint(*ptr)
	case *float64:
		return fmt.Sprint(*ptr)
	}
	panic(fmt.Sprintf("unsupported config field type %v", f.field.Type))
}

// set sets the value of field f in c to value.
func (c *config) set(f configField, value string) error {
	switch ptr := c.fieldPtr(f).(type) {
	case *string:
		if len(f.choices) > 0 {
			// Verify that v is one of the allowed choices.
			for _, choice := range f.choices {
				if choice == value {
					*ptr = value
					return nil
				}
			}
			return fmt.Errorf("invalid %q value %q", f.name, value)
		}
		*ptr = value
	case *bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		*ptr = v
	case *int:
		v, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		*ptr = v
	case *float64:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		*ptr = v
	default:
		panic(fmt.Sprintf("unsupported config field type %v", f.field.Type))
	}
	return nil
}

// defaultConfig returns the default configuration values.
// It is not affected by flags and interactive assignments.
func defaultConfig() config {
	return config{
		Endpoint:    "chat",
		MaxTokens:   math.MaxInt32,
		Temperature: 1,
	}
}

// isBoolConfig returns true if name is either the name of a boolean config field,
// or a valid value for a multi-choice config field.
func isBoolConfig(name string) bool {
	f, ok := configFieldMap[name]
	if !ok {
		return false
	}
	if name != f.name {
		return true
	}
	var c config
	_, ok = c.fieldPtr(f).(*bool)
	return ok
}

// isConfigurable returns true if name is either the name of a config field,
// or a valid value for a multi-choice config field.
func isConfigurable(name string) bool {
	_, ok := configFieldMap[name]
	return ok
}

// configure stores the name=value mapping into the current config,
// correctly handling the case when name identifies a particular
// choice in a field.
func configure(name, value string) error {
	currentCfgMu.Lock()
	defer currentCfgMu.Unlock()
	f, ok := configFieldMap[name]
	if !ok {
		return fmt.Errorf("unknown config field %q", name)
	}
	if f.name == name {
		return currentCfg.set(f, value)
	}
	// name must be one of the choices. If value is true,
	// set field-value to name.
	if v, err := strconv.ParseBool(value); v && err == nil {
		return currentCfg.set(f, name)
	}
	return fmt.Errorf("unknown config field %q", name)
}

// currentCfg holds the current configuration values.
// It is affected by flags and interactive assignments.
var currentCfg = defaultConfig()
var currentCfgMu sync.Mutex

func currentConfig() config {
	currentCfgMu.Lock()
	defer currentCfgMu.Unlock()
	return currentCfg
}

// configField contains metadata for a single configuration field.
type configField struct {
	name         string              // JSON field name/key in variables.
	urlParam     string              // URL parameter name.
	saved        bool                // Is field saved in settings?
	field        reflect.StructField // Field in config.
	choices      []string            // Name of variables in group.
	defaultValue string              // Default value for this field.
}

var (
	configFields []configField // Precomputed metadata per config field.

	// configFieldMap holds an entry for every config field as well as
	// an entry for every valid choice for fields with more than one choice.
	configFieldMap map[string]configField
)

func init() {
	// Config names for fields that are NOT saved in settings and
	// therefore do NOT have a JSON name.
	notSaved := map[string]string{}

	// choices holds the list of allowed values for config fields that
	// can take on one of a bounded set of values.
	choices := map[string][]string{
		"endpoint": {"chat", "completions", "models"},
	}

	// urlParam holds the mapping from a config field name to the URL
	// parameter used to hold that config field. If no entry is present
	// for a name, the corresponding field is not saved in URLs.
	urlParam := map[string]string{
		"endpoint":    "endpoint",
		"max_tokens":  "maxtokens",
		"temperature": "temp",
	}

	d := defaultConfig()
	configFieldMap = map[string]configField{}
	t := reflect.TypeOf(config{})
	for i, n := 0, t.NumField(); i < n; i++ {
		field := t.Field(i)
		json := strings.Split(field.Tag.Get("json"), ",")
		if len(json) == 0 {
			continue
		}
		// Get the configuration name for this field.
		name := json[0]
		if name == "-" {
			name = notSaved[field.Name]
			if name == "" {
				// Not a configuration field.
				continue
			}
		}
		f := configField{
			name:     name,
			urlParam: urlParam[name],
			saved:    name == json[0],
			field:    field,
			choices:  choices[name],
		}
		f.defaultValue = d.get(f)
		configFields = append(configFields, f)
		configFieldMap[f.name] = f
		for _, choice := range f.choices {
			configFieldMap[choice] = f
		}
	}
}

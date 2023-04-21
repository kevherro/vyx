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

// Package models implements the Models OpenAI endpoint.
package models

const (
	_ = "GET"
	_ = "https://api.openai.com/v1/models"
)

type Request struct{}

type Call interface {
	ParseTokens(input []string) ([]string, error)
}

func (r Request) ParseTokens(input []string) ([]string, error) {
	return nil, nil
}

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

// Package completions implements the Completions OpenAI endpoint.
package completions

const (
	_ = "POST"
	_ = "https://api.openai.com/v1/completions"
)

type Request struct {
	// ID of the model to use.
	Model string `json:"model"`

	// The prompt(s) to generate completions for.
	Prompt string `json:"prompt"`

	// The maximum number of tokens to generate in the completion.
	// Defaults to 16.
	MaxTokens int `json:"max_tokens"`

	// What sampling temperature to use, between 0 and 2.
	Temperature float64 `json:"temperature"`
}

type Response struct {
	ID           string       `json:"id"`
	Object       string       `json:"object"`
	CreatedAt    int64        `json:"created_at"`
	Model        string       `json:"model"`
	Choices      []Choice     `json:"choices"`
	Completion   Completion   `json:"completion"`
	Conversation Conversation `json:"conversation"`
}

type Completion struct {
	ID        string   `json:"id"`
	CreatedAt int64    `json:"created_at"`
	Model     string   `json:"model"`
	Prompt    string   `json:"prompt"`
	Choices   []Choice `json:"choices"`
}

type Choice struct {
	Text    string  `json:"text"`
	Index   int     `json:"index"`
	LogProb float64 `json:"logproba"`
}

type Conversation struct {
	ID        string    `json:"id"`
	CreatedAt int64     `json:"created_at"`
	Object    string    `json:"object"`
	Messages  []Message `json:"messages"`
}

type Message struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	Type      string `json:"type"`
	Author    string `json:"author"`
	Body      string `json:"body"`
}

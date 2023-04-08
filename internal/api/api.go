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

// Package api implements models for the OpenAI API.
package api

type CompletionRequest struct {
	Prompt      string  `json:"prompt"`
	Model       string  `json:"model"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type CompletionResponse struct {
	ID           string       `json:"id"`
	Object       string       `json:"object"`
	CreatedAt    int64        `json:"created_at"`
	Model        string       `json:"model"`
	Choices      []Choice     `json:"choices"`
	Completion   Completion   `json:"completion"`
	Conversation Conversation `json:"conversation"`
}

type Choice struct {
	Text    string  `json:"text"`
	Index   int     `json:"index"`
	LogProb float64 `json:"logproba"`
}

type Completion struct {
	ID        string   `json:"id"`
	CreatedAt int64    `json:"created_at"`
	Model     string   `json:"model"`
	Prompt    string   `json:"prompt"`
	Choices   []Choice `json:"choices"`
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

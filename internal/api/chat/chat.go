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

// Package chat implements the Chat OpenAI endpoint.
package chat

const (
	_ = "POST"
	_ = "https://api.openai.com/v1/chat/completions"
)

type Request struct {
	// ID of the model to use.
	Model string `json:"model"`

	// A list of messages describing the conversation so far.
	Messages []Message `json:"messages"`

	// The maximum number of tokens to generate in the completion.
	// Defaults to 16.
	MaxTokens int `json:"max_tokens"`

	// What sampling temperature to use, between 0 and 2.
	Temperature float64 `json:"temperature"`
}

type Message struct {
	// The role of the author of this message. One of system, user,
	// or assistant.
	Role string `json:"role"`

	// The contents of the message.
	Content string `json:"content"`
}

type Response struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

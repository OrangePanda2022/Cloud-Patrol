package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	http.HandleFunc("/chat", chatHandler)
	const PORT = "2121"
	log.Println("🚀 Server started at http://localhost:"+PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	// ===== CORS =====
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	// 处理预检请求
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// ===== SSE headers =====
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	const systemPrompt = "首先不要使用Markdown格式输出。其次你是一个帮助大学生，让大学生的生活变得方便的校园助手，使用日常说话的口吻，并且提供有用的建议"
	const accessToken = ""

	// ===== Client 创建 =====
	cfg := openai.DefaultConfig(accessToken)
	cfg.BaseURL = "https://aistudio.baidu.com/llm/lmapi/v3"

	client := openai.NewClientWithConfig(cfg)

	question := r.URL.Query().Get("question")
	if question == "" {
		question = "给我一些美好度过今天的建议"
	}

	req := openai.ChatCompletionRequest{
		Model: "ernie-4.5-0.3b",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    "user",
				Content: question,
			},
		},
		Stream:      true,
		MaxTokens:   8000,
		Temperature: 0.8,
		TopP:        0.8,
	}

	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		fmt.Fprintf(w, "event:error\ndata:%s\n\n", err.Error())
		flusher.Flush()
		return
	}
	defer stream.Close()

	for {
		resp, err := stream.Recv()
		if err != nil {
			fmt.Fprintf(w, "event:done\ndata:[DONE]\n\n")
			flusher.Flush()
			return
		}

		if len(resp.Choices) > 0 {
			delta := resp.Choices[0].Delta

			if delta.ReasoningContent != "" {
				fmt.Fprintf(w, "data:%s\n\n", delta.ReasoningContent)
			} else if delta.Content != "" {
				fmt.Fprintf(w, "data:%s\n\n", delta.Content)
			}

			flusher.Flush()
		}
	}
}

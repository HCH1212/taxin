package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/HCH1212/taxin/config"
	"github.com/pgvector/pgvector-go"
)

// 使用本地ollama all-minilm-l6-v2模型，仅支持英文，768维向量

// GenerateEmbeddingForLikes 根据文本生成词嵌入向量
func GenerateEmbeddingForLikes(ctx context.Context, likes []string) (pgvector.Vector, error) {
	var allEmbeddings []float32
	var baseURL string
	var model string

	if os.Getenv("GO_ENV") == "" {
		baseURL = "http://127.0.0.1:11434/api/embeddings"
		model = "chroma/all-minilm-l6-v2-f32:latest"
	} else {
		baseURL = config.GetConf().Ollama.Address
		model = config.GetConf().Ollama.Model
	}

	for _, like := range likes {
		// 构建请求体
		requestBody := map[string]string{
			"model":  model,
			"prompt": like,
		}
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			return pgvector.NewVector(nil), err
		}

		// 创建 HTTP 请求
		req, err := http.NewRequestWithContext(ctx, "POST", baseURL, bytes.NewBuffer(jsonBody))
		if err != nil {
			return pgvector.NewVector(nil), err
		}
		req.Header.Set("Content-Type", "application/json")

		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return pgvector.NewVector(nil), err
		}
		defer resp.Body.Close()

		// 读取响应
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return pgvector.NewVector(nil), err
		}

		// 检查响应状态码
		if resp.StatusCode != http.StatusOK {
			return pgvector.NewVector(nil), fmt.Errorf("request failed with status code: %d, body: %s", resp.StatusCode, string(body))
		}

		// 解析响应
		var result struct {
			Embedding []float32 `json:"embedding"`
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return pgvector.NewVector(nil), err
		}

		// 将当前喜好的嵌入向量追加到总向量中
		allEmbeddings = append(allEmbeddings, result.Embedding...)
	}

	return pgvector.NewVector(allEmbeddings), nil
}

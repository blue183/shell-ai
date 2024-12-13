package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var aiCmd = &cobra.Command{
	Use:   "cmd",
	Short: "调用大模型接口",
	Long:  `通过此命令调用大模型接口并获取响应`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("请提供输入文本")
			return
		}
		inputText := args[0]
		apiKey := "YOUR_API_KEY"                                             // 替换为您的 API 密钥
		url := "https://api.openai.com/v1/engines/davinci-codex/completions" // 替换为实际的 API URL

		// 创建请求体
		requestBody, err := json.Marshal(map[string]interface{}{
			"prompt":     inputText,
			"max_tokens": 100,
		})
		if err != nil {
			fmt.Println("请求体创建失败:", err)
			return
		}

		// 发送请求
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("请求创建失败:", err)
			return
		}
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("请求失败:", err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Println("请求失败，状态码:", response.StatusCode)
			return
		}

		var result map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			fmt.Println("解析响应失败:", err)
			return
		}

		// 打印 AI 响应
		if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
			if text, ok := choices[0].(map[string]interface{})["text"].(string); ok {
				fmt.Println("AI 响应:", text)
			}
		}
	},
}

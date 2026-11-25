package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/voocel/litellm"
	"github.com/voocel/litellm/provider/openai"

	"vadik/core"
	"vadik/tools"
)

var (
	configPath string
)

func main() {
	flag.StringVar(&configPath, "config", "config.json", "config file path in json")

	flag.Parse()
	// Loading OpenAI API data
	cfg := core.ConfigFromFile(configPath)

	// Initialization of agent
	provider, err := openai.New(
		openai.Config{
			BaseURL: cfg.OpenAIAPIBase,
			APIKey:  cfg.OpenAIAPIToken,
		},
	)
	if err != nil {
		log.Fatalf("Failed to init provider: %v\n", err)
	}

	client, err := litellm.New(provider)
	if err != nil {
		log.Fatalf("Failed to init client: %v\n", err)
	}

	ctx := context.Background()

	messages := []litellm.Message{
		litellm.System(cfg.SystemPrompt),
	}

	// Main loop
	scanner := bufio.NewScanner(os.Stdin)
	prompt := ""
	lastToolCallID := 0
	wasToolCalled := false
	for {
		if !wasToolCalled {
			fmt.Print("\n[Prompt] << ")
			if !scanner.Scan() {
				fmt.Println("Bye ...")
				break
			}
			prompt = scanner.Text()
			if len(prompt) <= 0 {
				continue
			}
			messages = append(messages, litellm.UserText(prompt))
		}

		// Generating response
		resp, err := core.Stream(ctx, client, litellm.Request{
			Model:      cfg.Model,
			Messages:   messages,
			Tools:      tools.Tools,
			ToolChoice: &litellm.ToolChoice{Mode: "auto"},
		})
		fmt.Println()
		if err != nil {
			log.Fatalf("Failed to ask AI: %v\n", err)
		}
		messages = append(messages, litellm.Assistant(resp.Blocks...))
		toolCalls := resp.ToolCalls()
		if len(toolCalls) <= 0 {
			wasToolCalled = false
			continue
		}

		wasToolCalled = true
		for _, tool := range toolCalls {
			fmt.Printf("\n[TOOL] `%v` with args: %v\n", tool.Name, tool.Arguments)
			toolFunc, exists := tools.ToolsMap[tool.Name]
			toolIdName := fmt.Sprintf("call%v", lastToolCallID)
			lastToolCallID++
			if !exists {
				messages = append(messages, litellm.ToolResultText(toolIdName, "tool not found!"))
				continue
			}
			result := toolFunc(*cfg, tool)
			fmt.Printf("\n[CALL] %v: %v\n", toolIdName, result)
			messages = append(messages, litellm.ToolResultText(toolIdName, result))
		}
	}

}

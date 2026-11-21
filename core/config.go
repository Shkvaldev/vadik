package core

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// VadikConfig is go representation of config
type VadikConfig struct {
	OpenAIAPIBase  string `json:"openai_api_base"`
	OpenAIAPIToken string `json:"openai_api_token"`
	Model          string `json:"model"`
	SystemPrompt   string `json: "system_prompt"`
}

func (c VadikConfig) String() string {
	return fmt.Sprintf("Config [base=%v, model=%v]", c.OpenAIAPIBase, c.Model)
}

// ConfigFromFile reads json cfg from file, returns config struct
func ConfigFromFile(path string) *VadikConfig {
	cfg := &VadikConfig{}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %v\n", err)
	}
	if err = json.Unmarshal(data, cfg); err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	if len(cfg.OpenAIAPIBase) <= 0 || len(cfg.Model) <= 0 {
		log.Fatalf("Failed to load config: api base and model must not be empty!")
	}

	if len(cfg.SystemPrompt) <= 0 {
		cfg.SystemPrompt = "You are a helpful assistant. Answer shortly, ask questions and use tools if required."
	}

	return cfg
}

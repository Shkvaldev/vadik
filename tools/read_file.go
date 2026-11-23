package tools

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/voocel/litellm"
)

// Read file tool rule
var ReadFileRule, _ = litellm.NewTool(
	"read_file",
	"Read file.",
	map[string]any{
		"type": "object",
		"properties": map[string]any{
			"path": map[string]any{"type": "string"},
		},
		"required": []string{"path"},
	})

// Read file tool function
func ReadFileTool(call litellm.ToolUseBlock) string {
	var args struct {
		Path string `json:"path"`
	}
	_ = json.Unmarshal([]byte(call.Arguments), &args)
	if len(args.Path) <= 0 {
		return "no path provided"
	}
	content, err := os.ReadFile(args.Path)
	if err != nil {
		return fmt.Sprintf("Can not read file: %v", err)
	}
	return fmt.Sprintf("Output: %v", string(content))
}

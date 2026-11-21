package tools

import (
	"encoding/json"
	"fmt"

	"github.com/voocel/litellm"
)

// Shell cmd tool rule
var RunShellRule, _ = litellm.NewTool(
	"run_shell",
	"Execute bash command.",
	map[string]any{
		"type": "object",
		"properties": map[string]any{
			"command": map[string]any{"type": "string"},
		},
		"required": []string{"command"},
	})

// Shell cmd tool function
func RunShellTool(call litellm.ToolUseBlock) string {
	var args struct {
		Command string `json: "command"`
	}
	_ = json.Unmarshal([]byte(call.Arguments), &args)
	if len(args.Command) <= 0 {
		return fmt.Sprintf("no command provided")
	}
	return "not implemented"
}

package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

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
			"timeout": map[string]any{"type": "integer"},
		},
		"required": []string{"command"},
	})

// Shell cmd tool function
// BUG: THERE IS NO SECURITY!!! ADD AS SOON AS POSSIBLE!!!
func RunShellTool(call litellm.ToolUseBlock) string {
	var args struct {
		Command string `json:"command"`
		Timeout int    `json:"timeout"`
	}
	_ = json.Unmarshal([]byte(call.Arguments), &args)
	if len(args.Command) <= 0 {
		return "no command provided"
	}

	// Default timeout for cmd is 5 secs
	timeout := 5 * time.Second
	if args.Timeout > 0 {
		timeout = time.Duration(args.Timeout) * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sh", "-c", args.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("could not run command: %v %s", err, out)
	}
	return fmt.Sprintf("Output: %v", string(out))
}

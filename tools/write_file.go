package tools

import (
	"encoding/json"
	"fmt"
	"github.com/voocel/litellm"
	"os"
	"vadik/core"
)

// Write file tool rule
var WriteFileRule, _ = litellm.NewTool(
	"write_file",
	"Write data to file.",
	map[string]any{
		"type": "object",
		"properties": map[string]any{
			"path": map[string]any{"type": "string"},
			"data": map[string]any{"type": "string"},
		},
		"required": []string{"path", "data"},
	})

// Write file tool function
func WriteFileTool(cfg core.VadikConfig, call litellm.ToolUseBlock) string {
	var args struct {
		Path string `json:"path"`
		Data string `json:"data"`
	}
	_ = json.Unmarshal([]byte(call.Arguments), &args)
	if len(args.Path) <= 0 {
		return "no path provided"
	}
	if len(args.Data) <= 0 {
		return "no data provided"
	}

	// Security policy
	if cfg.SecurityLevel == 1 {
		if !StdAskCallback(fmt.Sprintf("write to file `%s`?", args.Path)) {
			return fmt.Sprintf("writing to `%s` denied by user", args.Path)
		}
	}

	err := os.WriteFile(args.Path, []byte(args.Data), 0644)
	if err != nil {
		return fmt.Sprintf("failed to write to file `%v`: %v", args.Path, err)
	}
	return fmt.Sprintf("written successfully to `%v`", args.Path)
}

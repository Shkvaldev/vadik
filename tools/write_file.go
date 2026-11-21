package tools

import (
        "encoding/json"
        "fmt"

        "github.com/voocel/litellm"
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
func WriteFileTool(call litellm.ToolUseBlock) string {
        var args struct {
                Path string `json:"path"`
                Data string `json:"data"`
        }
        _ = json.Unmarshal([]byte(call.Arguments), &args)
        if len(args.Path) <= 0 {
                return fmt.Sprintf("no path provided")
        }
        if len(args.Data) <= 0 {
                return fmt.Sprintf("no data provided")
        }
        return "not implemented"
}

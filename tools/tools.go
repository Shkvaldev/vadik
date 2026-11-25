package tools

import (
	"github.com/voocel/litellm"
	"vadik/core"
)

// Tools list
var Tools []litellm.Tool = []litellm.Tool{
	RunShellRule,
	ReadFileRule,
	WriteFileRule,
}

type ToolFunc func(cfg core.VadikConfig, call litellm.ToolUseBlock) string

var ToolsMap map[string]ToolFunc = map[string]ToolFunc{
	RunShellRule.Name:  RunShellTool,
	ReadFileRule.Name:  ReadFileTool,
	WriteFileRule.Name: WriteFileTool,
}

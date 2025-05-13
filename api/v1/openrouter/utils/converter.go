package utils

import (
	"crynux_bridge/api/v1/openrouter/structs"
)

func ChatCompletionsRoleToRole(role structs.ChatCompletionsRole) structs.Role {
	switch role {
	case structs.ChatCompletionsRoleDeveloper:
		return structs.RoleUnknown
	case structs.ChatCompletionsRoleSystem:
		return structs.RoleSystem
	case structs.ChatCompletionsRoleUser:
		return structs.RoleUser
	case structs.ChatCompletionsRoleAssistant:
		return structs.RoleAssistant
	case structs.ChatCompletionsRoleTool:
		return structs.RoleTool
	}
	return structs.RoleUnknown
}

func RoleToChatCompletionsRole(role structs.Role) structs.ChatCompletionsRole {
	switch role {
	case structs.RoleUnknown:
		return structs.ChatCompletionsRoleUnknown
	case structs.RoleSystem:
		return structs.ChatCompletionsRoleSystem
	case structs.RoleUser:
		return structs.ChatCompletionsRoleUser
	case structs.RoleAssistant:
		return structs.ChatCompletionsRoleAssistant
	case structs.RoleTool:
		return structs.ChatCompletionsRoleTool
	}
	return structs.ChatCompletionsRoleUnknown
}

func CCReqMessageToolCallToToolCall(ccrMessagetoolCall structs.CCReqMessageToolCall) map[string]interface{} {
	toolCall := make(map[string]interface{})
	toolCall["id"] = ccrMessagetoolCall.ID
	toolCall["type"] = ccrMessagetoolCall.Type

	function := make(map[string]string)
	function["name"] = ccrMessagetoolCall.Function.Name
	function["arguments"] = ccrMessagetoolCall.Function.Arguments

	toolCall["function"] = function
	return toolCall
}

func CCReqMessageToMessage(ccrMessage structs.CCReqMessage) structs.Message {
	var message structs.Message
	message.Role = ChatCompletionsRoleToRole(ccrMessage.Role)
	message.Content = ccrMessage.Content
	message.ToolCallID = ccrMessage.ToolCallID

	toolCalls := make([]map[string]interface{}, len(ccrMessage.ToolCalls))
	for i, toolCall := range ccrMessage.ToolCalls {
		toolCalls[i] = CCReqMessageToolCallToToolCall(toolCall)
	}
	message.ToolCalls = toolCalls

	return message
}

func MessageToCCResMessage(message structs.Message) structs.CCResMessage {
	var ccResMessage structs.CCResMessage
	ccResMessage.Role = RoleToChatCompletionsRole(message.Role)
	ccResMessage.Content = message.Content
	// ccResMessage.Refusal = ""
	// ccResMessage.Annotations = nil
	// ccResMessage.Audio = nil
	ccResMessage.ToolCalls = message.ToolCalls

	return ccResMessage
}

func ResponseChoiceToCCResChoice(responseChoice structs.ResponseChoice) structs.CCResChoice {
	var ccResChoice structs.CCResChoice
	ccResChoice.Index = responseChoice.Index
	ccResChoice.Message = MessageToCCResMessage(responseChoice.Message)
	ccResChoice.LogProbs = nil
	ccResChoice.FinishReason = string(responseChoice.FinishReason)
	return ccResChoice
}

func UsageToCCResUsage(usage structs.Usage) structs.CCResUsage {
	var ccResUsage structs.CCResUsage
	ccResUsage.PromptTokens = usage.PromptTokens
	ccResUsage.CompletionTokens = usage.CompletionTokens
	ccResUsage.TotalTokens = usage.TotalTokens
	// ccResUsage.PromptTokensDetails = structs.PromptTokensDetails{}
	// ccResUsage.CompletionTokensDetails = structs.CompletionTokensDetails{}
	return ccResUsage
}

func ResponseChoiceToCResChoice(responseChoice structs.ResponseChoice) (structs.CResChoice, error) {
	var cResChoice structs.CResChoice
	cResChoice.Index = responseChoice.Index
	cResChoice.Text = responseChoice.Message.Content
	// ccResChoice.LogProbs = ""
	cResChoice.FinishReason = string(responseChoice.FinishReason)
	return cResChoice, nil
}

func UsageToCResUsage(usage structs.Usage) structs.CResUsage {
	var cResUsage structs.CResUsage
	cResUsage.PromptTokens = usage.PromptTokens
	cResUsage.CompletionTokens = usage.CompletionTokens
	cResUsage.TotalTokens = usage.TotalTokens
	return cResUsage
}

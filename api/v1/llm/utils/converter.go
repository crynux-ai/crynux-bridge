package utils

import (
	"crynux_bridge/api/v1/llm/structs"
	"crynux_bridge/models"
)

func ChatCompletionsRoleToRole(role structs.ChatCompletionsRole) models.LLMRole {
	switch role {
	case structs.ChatCompletionsRoleDeveloper:
		return models.LLMRoleUnknown
	case structs.ChatCompletionsRoleSystem:
		return models.LLMRoleSystem
	case structs.ChatCompletionsRoleUser:
		return models.LLMRoleUser
	case structs.ChatCompletionsRoleAssistant:
		return models.LLMRoleAssistant
	case structs.ChatCompletionsRoleTool:
		return models.LLMRoleTool
	}
	return models.LLMRoleUnknown
}

func RoleToChatCompletionsRole(role models.LLMRole) structs.ChatCompletionsRole {
	switch role {
	case models.LLMRoleUnknown:
		return structs.ChatCompletionsRoleUnknown
	case models.LLMRoleSystem:
		return structs.ChatCompletionsRoleSystem
	case models.LLMRoleUser:
		return structs.ChatCompletionsRoleUser
	case models.LLMRoleAssistant:
		return structs.ChatCompletionsRoleAssistant
	case models.LLMRoleTool:
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

func CCReqMessageToMessage(ccrMessage structs.CCReqMessage) models.Message {
	var message models.Message
	message.Role = ChatCompletionsRoleToRole(ccrMessage.Role)
	message.Content = ccrMessage.Content
	message.ToolCallID = ccrMessage.ToolCallID

	if len(ccrMessage.ToolCalls) > 0 {
		message.ToolCalls = make([]structs.ToolCall, len(ccrMessage.ToolCalls))
		for i, reqToolCall := range ccrMessage.ToolCalls {
			message.ToolCalls[i] = structs.ToolCall{
				Id:   reqToolCall.ID,
				Type: reqToolCall.Type,
				Function: structs.FunctionCall{
					Name:      reqToolCall.Function.Name,
					Arguments: reqToolCall.Function.Arguments,
				},
			}
		}
	}

	return message
}

func MessageToCCResMessage(message models.Message) structs.CCResMessage {
	var ccResMessage structs.CCResMessage
	ccResMessage.Role = RoleToChatCompletionsRole(message.Role)
	ccResMessage.Content = message.Content
	// ccResMessage.Refusal = ""
	// ccResMessage.Annotations = nil
	// ccResMessage.Audio = nil
	ccResMessage.ToolCalls = message.ToolCalls

	return ccResMessage
}

func ResponseChoiceToCCResChoice(responseChoice models.ResponseChoice) structs.CCResChoice {
	var ccResChoice structs.CCResChoice
	ccResChoice.Index = responseChoice.Index
	ccResChoice.Message = MessageToCCResMessage(responseChoice.Message)
	ccResChoice.LogProbs = nil
	ccResChoice.FinishReason = string(responseChoice.FinishReason)
	return ccResChoice
}

func UsageToCCResUsage(usage models.Usage) structs.CCResUsage {
	var ccResUsage structs.CCResUsage
	ccResUsage.PromptTokens = usage.PromptTokens
	ccResUsage.CompletionTokens = usage.CompletionTokens
	ccResUsage.TotalTokens = usage.TotalTokens
	// ccResUsage.PromptTokensDetails = structs.PromptTokensDetails{}
	// ccResUsage.CompletionTokensDetails = structs.CompletionTokensDetails{}
	return ccResUsage
}

func ResponseChoiceToCResChoice(responseChoice models.ResponseChoice) (structs.CResChoice, error) {
	var cResChoice structs.CResChoice
	cResChoice.Index = responseChoice.Index
	cResChoice.Text = responseChoice.Message.Content
	// ccResChoice.LogProbs = ""
	cResChoice.FinishReason = string(responseChoice.FinishReason)
	return cResChoice, nil
}

func UsageToCResUsage(usage models.Usage) structs.CResUsage {
	var cResUsage structs.CResUsage
	cResUsage.PromptTokens = usage.PromptTokens
	cResUsage.CompletionTokens = usage.CompletionTokens
	cResUsage.TotalTokens = usage.TotalTokens
	return cResUsage
}

package tests

const GPTResponseStr = `{
	"model": "gpt2",
	"choices": [
		{
			"finish_reason": "length",
			"message": {
				"role": "assistant",
				"content": '\n\nI have a chat bot, called "Eleanor" which was developed by my team on Skype. '
				"The only thing I will say is this",
			},
			"index": 0,
		}
	],
	"usage": {"prompt_tokens": 11, "completion_tokens": 30, "total_tokens": 41},
}`

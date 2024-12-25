package main

import (
	"context"
  "encoding/json"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/shared"
)

var surveyTools = []openai.ChatCompletionToolParam{
  {
    Type: openai.F(openai.ChatCompletionToolTypeFunction),
    Function: openai.F(shared.FunctionDefinitionParam{
      Name: openai.F("ask"),
      Description: openai.F("Ask a question"),
      Parameters: openai.F(shared.FunctionParameters{
        "type": "object",
        "properties": map[string]any{
          "prompt": map[string]any{
            "type": "string",
          },
        },
      }),
    }),
  },
  {
    Type: openai.F(openai.ChatCompletionToolTypeFunction),
    Function: openai.F(shared.FunctionDefinitionParam{
      Name: openai.F("confirm"),
      Description: openai.F("Ask a yes/no question"),
      Parameters: openai.F(shared.FunctionParameters{
        "type": "object",
        "properties": map[string]any{
          "prompt": map[string]any{
            "type": "string",
          },
        },
      }),
    }),
  },
  {
    Type: openai.F(openai.ChatCompletionToolTypeFunction),
    Function: openai.F(shared.FunctionDefinitionParam{
      Name: openai.F("select"),
      Description: openai.F("Ask a multiple choice question"),
      Parameters: openai.F(shared.FunctionParameters{
        "type": "object",
        "properties": map[string]any{
          "prompt": map[string]any{
            "type": "string",
          },
          "options": map[string]any{
            "type": "array",
            "items": map[string]any{
              "type": "string",
            },
          },
        },
      }),
    }),
  },
}

func main() {
  ctx := context.Background()
  client := openai.NewClient()
  history := []openai.ChatCompletionMessageParamUnion{
      openai.SystemMessage(
        `Your job is to survey the user about their use of the application they just used, "Smart SaaS".`+
        ` Utilize the provided tools to ask questions and get structured responses.`+
        ` Only use the tools to communicate with the user.`+
        ` If you're asking a question that requires a specific type of response, use the appropriate tool.`+
        ` For example, if you're asking a yes/no question, use the "confirm" tool or if it can be a multiple choice, use "select".`,
      ),
  }

  com, err := prompt(ctx, client, history)
  if err != nil {
    panic(err)
  }

  b, err := json.MarshalIndent(com, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Printf("Response:\n%s\n", string(b))
}

func prompt(ctx context.Context, client *openai.Client, hist []openai.ChatCompletionMessageParamUnion) (*openai.ChatCompletion, error) {
  return client.Chat.Completions.New(
    ctx,
    openai.ChatCompletionNewParams{
      Model: openai.F("gpt-4o-mini"),
      N: openai.F[int64](1),
      Messages: openai.F(hist),
      Tools: openai.F(surveyTools),
    },
  )
}

type askprompt struct {
  Prompt string `json:"prompt"`
}

type confirmprompt struct {
  Prompt string `json:"prompt"`
}

type selectprompt struct {
  Prompt string `json:"prompt"`
  Options []string `json:"options"`
}

func (sp selectprompt) ask() (string, error) {

}

func parseToolCall(tc openai.ChatCompletionToolCall) 

func handleResponse(com *openai.ChatCompletion) {
  ch := com.Choices[0]
  if len(ch.Message.Content) > 0 {
    fmt.Println("> " + ch.Message.Content)
  }
  for _, tc := range ch.Message.ToolCalls {
    fn := tc.Function
  }
  for _, msg := range com.Choices[0].Messages {
    fmt.Println(msg.Content)
  }
}


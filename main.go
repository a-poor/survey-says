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
  com, err := client.Chat.Completions.New(
    ctx,
    openai.ChatCompletionNewParams{
      Model: openai.F("gpt-4o-mini"),
      N: openai.F[int64](1),
      Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
        openai.SystemMessage(
          `Your job is to survey the user about their use of the application they just used, "Smart SaaS".`+
          ` Utilize the provided tools to ask questions and get structured responses.`+
          ` If you're asking a question that requires a specific type of response, use the appropriate tool.`+
          ` For example, if you're asking a yes/no question, use the "confirm" tool or if it can be a multiple choice, use "select".`,
        ),
      }),
      Tools: openai.F(surveyTools),
    },
  )
  if err != nil {
    panic(err)
  }

  b, err := json.MarshalIndent(com, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Printf("Response:\n%s\n", string(b))
}



package gemini

import "github.com/xxxsen/common/replacer"

type Request struct {
	Contents []Content `json:"contents"`
}

type Response struct {
	Candidates    []Candidate   `json:"candidates"`
	UsageMetadata UsageMetadata `json:"usageMetadata"`
	ModelVersion  string        `json:"modelVersion"`
}

type Candidate struct {
	Content      Content `json:"content"`
	FinishReason string  `json:"finishReason"`
	AvgLogprobs  float64 `json:"avgLogprobs"`
}

type Content struct {
	Parts []Part `json:"parts"`
	Role  string `json:"role"`
}

type Part struct {
	Text string `json:"text"`
}

type UsageMetadata struct {
	PromptTokenCount        int           `json:"promptTokenCount"`
	CandidatesTokenCount    int           `json:"candidatesTokenCount"`
	TotalTokenCount         int           `json:"totalTokenCount"`
	PromptTokensDetails     []TokenDetail `json:"promptTokensDetails"`
	CandidatesTokensDetails []TokenDetail `json:"candidatesTokensDetails"`
}

type TokenDetail struct {
	Modality   string `json:"modality"`
	TokenCount int    `json:"tokenCount"`
}

func buildRequest(prompt string, m map[string]interface{}) *Request {
	res := replacer.ReplaceByMap(prompt, m)
	content := Content{
		Parts: []Part{
			{
				Text: res,
			},
		},
	}
	return &Request{
		Contents: []Content{content},
	}
}

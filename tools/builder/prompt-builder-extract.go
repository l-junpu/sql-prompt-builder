package builder

import "strings"

type DataExtractor interface {
	Generate() string
}

func (b *PromptBuilder) Generate() string {
	b.Prompt = strings.TrimSpace(b.Prompt)
	return b.Prompt + ";"
}

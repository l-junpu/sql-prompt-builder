package builder

import (
	"fmt"
	"strings"
)

type DataLimiters interface {
	Limit(Rows int) PromptBuilderInterface
	Offset(Rows int) PromptBuilderInterface
}

func (b *PromptBuilder) Limit(Rows int) PromptBuilderInterface {
	b.Prompt = fmt.Sprintf("%sLIMIT %v\n", b.Prompt, Rows)
	return b
}

func (b *PromptBuilder) Offset(Rows int) PromptBuilderInterface {
	trimmed_prompt := strings.TrimSpace(b.Prompt)
	ss := strings.Split(trimmed_prompt, "\n")

	if strings.Contains(ss[len(ss)-1], "LIMIT") {
		b.Prompt = strings.TrimSpace(b.Prompt)
		b.Prompt = fmt.Sprintf("%s OFFSET %v\n", b.Prompt, Rows)
	} else {
		b.Prompt = fmt.Sprintf("%sOFFSET %v\n", b.Prompt, Rows)
	}
	return b
}

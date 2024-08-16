package builder

import "fmt"

type DataGrouping interface {
	GroupBy(Columns ...string) PromptBuilderInterface
	Having() PromptBuilderInterface
	Rollup() PromptBuilderInterface
}

func (b *PromptBuilder) GroupBy(Columns ...string) PromptBuilderInterface {
	for i, col := range Columns {
		if i == 0 {
			b.Prompt = fmt.Sprintf("%sGROUP BY %s", b.Prompt, col)
		} else {
			b.Prompt = fmt.Sprintf("%s, %s", b.Prompt, col)
		}
	}
	b.Prompt = fmt.Sprintf("%s\n", b.Prompt)
	return b
}

func (b *PromptBuilder) Having() PromptBuilderInterface {
	b.Prompt = b.Prompt + " Having"
	return b
}

func (b *PromptBuilder) Rollup() PromptBuilderInterface {
	b.Prompt = b.Prompt + " Rollup"
	return b
}

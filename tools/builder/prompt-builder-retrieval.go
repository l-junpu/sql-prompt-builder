package builder

import "fmt"

type DataRetrieval interface {
	Select(Columns ...string) PromptBuilderInterface
	From(Table string) PromptBuilderInterface
	InnerJoin(FirstTable string, SecondTable string, FirstCol string, SecondCol string) PromptBuilderInterface
}

func (b *PromptBuilder) Select(Columns ...string) PromptBuilderInterface {
	for i, col := range Columns {
		if i == 0 {
			b.Prompt = fmt.Sprintf("SELECT %s", col)
		} else {
			b.Prompt = fmt.Sprintf("%s, %s", b.Prompt, col)
		}
	}
	b.Prompt = fmt.Sprintf("%s\n", b.Prompt)

	return b
}

func (b *PromptBuilder) From(Table string) PromptBuilderInterface {
	b.Prompt = fmt.Sprintf("%sFROM %s\n", b.Prompt, Table)
	return b
}

func (b *PromptBuilder) InnerJoin(FirstTable string, SecondTable string, FirstCol string, SecondCol string) PromptBuilderInterface {
	b.Prompt = fmt.Sprintf("%sINNER JOIN %s ON %s.%s = %s.%s\n", b.Prompt, SecondTable, FirstTable, FirstCol, SecondTable, SecondCol)
	return b
}

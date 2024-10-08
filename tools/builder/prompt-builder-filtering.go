package builder

import (
	"fmt"
	"strings"
)

type DataFilters interface {
	Where(Condition string, Comparison ...interface{}) PromptBuilderInterface
	And(Condition string, Comparison interface{}) PromptBuilderInterface
	Or(Condition string, Comparison interface{}) PromptBuilderInterface

	conditionalWrapBraces(Keyword string, Comparisons ...interface{}) PromptBuilderInterface
	In(Comparisons ...interface{}) PromptBuilderInterface
	NotIn(Comparisons ...interface{}) PromptBuilderInterface

	// Between() PromptBuilderInterface
	// NotBetween() PromptBuilderInterface

	Like(Pattern string) PromptBuilderInterface
	// NotLike() PromptBuilderInterface

	// IsNull() PromptBuilderInterface
	// IsNotNull() PromptBuilderInterface

	// Exists() PromptBuilderInterface
	// NotExists() PromptBuilderInterface
}

func (b *PromptBuilder) Where(Condition string, Comparison ...interface{}) PromptBuilderInterface {
	Condition = fmt.Sprintf(Condition, Comparison...)
	b.Prompt = fmt.Sprintf("%sWHERE %s\n", b.Prompt, Condition)
	return b
}

func (b *PromptBuilder) And(Condition string, Comparison interface{}) PromptBuilderInterface {
	b.Prompt = strings.TrimSpace(b.Prompt)
	Condition = fmt.Sprintf(Condition, Comparison)
	b.Prompt = fmt.Sprintf("%s AND %s\n", b.Prompt, Condition)
	return b
}

func (b *PromptBuilder) Or(Condition string, Comparison interface{}) PromptBuilderInterface {
	b.Prompt = strings.TrimSpace(b.Prompt)
	Condition = fmt.Sprintf(Condition, Comparison)
	b.Prompt = fmt.Sprintf("%s OR %s\n", b.Prompt, Condition)
	return b
}

func (b *PromptBuilder) conditionalWrapBraces(Keyword string, Comparisons ...interface{}) PromptBuilderInterface {
	b.Prompt = strings.TrimSpace(b.Prompt)
	if len(Comparisons) > 1 {
		concat := "("
		for _, c := range Comparisons {
			concat = fmt.Sprintf("%s'%s', ", concat, c)
		}
		concat = strings.TrimSuffix(concat, string(concat[len(concat)-2:]))
		concat += ")"
		b.Prompt = fmt.Sprintf("%s %s %s\b", b.Prompt, Keyword, concat)
	} else {
		b.Prompt = fmt.Sprintf("%s %s %s\n", b.Prompt, Keyword, Comparisons)
	}
	b.Prompt = fmt.Sprintf("%s\n", b.Prompt)
	return b
}

func (b *PromptBuilder) In(Comparisons ...interface{}) PromptBuilderInterface {
	return b.conditionalWrapBraces("IN", Comparisons...)
}

func (b *PromptBuilder) NotIn(Comparisons ...interface{}) PromptBuilderInterface {
	return b.conditionalWrapBraces("NOT IN", Comparisons...)
}

func (b *PromptBuilder) Like(Pattern string) PromptBuilderInterface {
	b.Prompt = strings.TrimSpace(b.Prompt)
	b.Prompt = fmt.Sprintf("%s LIKE %s\n", b.Prompt, Pattern)
	return b
}

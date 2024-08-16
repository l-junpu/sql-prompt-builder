package builder

type DataAggregation interface {
	Sum() PromptBuilderInterface
	Avg() PromptBuilderInterface
	Max() PromptBuilderInterface
	Min() PromptBuilderInterface
	Count() PromptBuilderInterface
}

func (b *PromptBuilder) Sum() PromptBuilderInterface {
	b.Prompt = b.Prompt + " Sum"
	return b
}

func (b *PromptBuilder) Avg() PromptBuilderInterface {
	b.Prompt = b.Prompt + " Avg"
	return b
}

func (b *PromptBuilder) Max() PromptBuilderInterface {
	b.Prompt = b.Prompt + " Max"
	return b
}

func (b *PromptBuilder) Min() PromptBuilderInterface {
	b.Prompt = b.Prompt + " Min"
	return b
}

func (b *PromptBuilder) Count() PromptBuilderInterface {
	b.Prompt = b.Prompt + " Count"
	return b
}

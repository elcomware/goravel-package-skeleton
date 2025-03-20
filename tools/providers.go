package tools

type ProviderTools struct {
	ProviderName *string
}

// PublishesServiceProvider sets the provider name as publishable.
func (p *ProviderTools) PublishesServiceProvider(providerName string) *ProviderTools {
	p.ProviderName = &providerName
	return p
}

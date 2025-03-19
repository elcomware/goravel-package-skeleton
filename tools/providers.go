package tools

type PackageProviders struct {
	ProviderName *string
}

// PublishesServiceProvider sets the provider name as publishable.
func (p *PackageProviders) PublishesServiceProvider(providerName string) *PackageProviders {
	p.ProviderName = &providerName
	return p
}

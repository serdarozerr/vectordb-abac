package auth

import (
	"github.com/serdarozerr/vectordb-abac/config"
	"net/url"
)

type Director struct {
	builder *Builder
}

type Builder struct {
	values *url.Values
}

func NewBuilder() *Builder {
	return &Builder{values: &url.Values{}}
}

func (b *Builder) SetAttribute(attribute string, value interface{}) *Builder {
	b.values.Set(attribute, value.(string))
	return b
}

func (b *Builder) Build() *url.Values {
	return b.values
}

func (d *Director) BuildAuthCodeData(conf *config.Config, code string, grantType string) *url.Values {
	d.builder.
		SetAttribute("code", code).
		SetAttribute("grant_type", grantType).
		SetAttribute("client_id", conf.Auth.ClientID).
		SetAttribute("client_secret", conf.Auth.ClientSecret).
		SetAttribute("redirect_uri", conf.Auth.RedirectURI)
	data := d.builder.Build()

	return data
}

func (d *Director) BuildRefreshTokenData(conf *config.Config, refresh_token string, grantType string) *url.Values {
	d.builder.
		SetAttribute("refresh_token", refresh_token).
		SetAttribute("grant_type", grantType).
		SetAttribute("client_id", conf.Auth.ClientID).
		SetAttribute("client_secret", conf.Auth.ClientSecret)
	data := d.builder.Build()

	return data
}

package turso

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "turso",
		Platform: schema.PlatformInfo{
			Name:     "Turso",
			Homepage: sdk.URL("https://turso.tech"),
		},
		Credentials: []schema.CredentialType{
			PersonalAccessToken(),
		},
		Executables: []schema.Executable{
			TursoCLI(),
		},
	}
}

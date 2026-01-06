package resources

import "github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"

var DeleteSecurityIntegration = ResourceDeleteContextFunc(
	sdk.ParseAccountObjectIdentifier,
	func(client *sdk.Client) DropSafelyFunc[sdk.AccountObjectIdentifier] {
		return client.SecurityIntegrations.DropSafely
	},
)

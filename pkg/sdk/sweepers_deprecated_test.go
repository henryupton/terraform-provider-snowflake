package sdk_test

import (
	"context"
	"log"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func getAccountPolicyAttachmentsSweeper(client *sdk.Client) func() error {
	return func() error {
		log.Printf("[DEBUG] Unsetting password and session policies set on the account level")
		ctx := context.Background()
		_ = client.Accounts.UnsetPolicySafely(ctx, sdk.PolicyKindPasswordPolicy)
		_ = client.Accounts.UnsetPolicySafely(ctx, sdk.PolicyKindSessionPolicy)
		return nil
	}
}

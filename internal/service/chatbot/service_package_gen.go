// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package chatbot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chatbot"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceSlackWorkspace,
			TypeName: "aws_chatbot_slack_workspace",
			Name:     "Slack Workspace",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{
		{
			Factory:  newSlackChannelConfigurationResource,
			TypeName: "aws_chatbot_slack_channel_configuration",
			Name:     "Slack Channel Configuration",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "chat_configuration_arn",
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newTeamsChannelConfigurationResource,
			TypeName: "aws_chatbot_teams_channel_configuration",
			Name:     "Teams Channel Configuration",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "chat_configuration_arn",
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Chatbot
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*chatbot.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*chatbot.Options){
		chatbot.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *chatbot.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "chatbot",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return chatbot.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*chatbot.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*chatbot.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *chatbot.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*chatbot.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

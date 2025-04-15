// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package bedrock

import (
	"context"
	"unique"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newCustomModelDataSource,
			TypeName: "aws_bedrock_custom_model",
			Name:     "Custom Model",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newCustomModelsDataSource,
			TypeName: "aws_bedrock_custom_models",
			Name:     "Custom Models",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newFoundationModelDataSource,
			TypeName: "aws_bedrock_foundation_model",
			Name:     "Foundation Model",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newFoundationModelsDataSource,
			TypeName: "aws_bedrock_foundation_models",
			Name:     "Foundation Models",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newInferenceProfileDataSource,
			TypeName: "aws_bedrock_inference_profile",
			Name:     "Inference Profile",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newInferenceProfilesDataSource,
			TypeName: "aws_bedrock_inference_profiles",
			Name:     "Inference Profiles",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{
		{
			Factory:  newCustomModelResource,
			TypeName: "aws_bedrock_custom_model",
			Name:     "Custom Model",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: "job_arn",
			}),
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newResourceGuardrail,
			TypeName: "aws_bedrock_guardrail",
			Name:     "Guardrail",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: "guardrail_arn",
			}),
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newGuardrailVersionResource,
			TypeName: "aws_bedrock_guardrail_version",
			Name:     "Guardrail Version",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newResourceInferenceProfile,
			TypeName: "aws_bedrock_inference_profile",
			Name:     "Inference Profile",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newModelInvocationLoggingConfigurationResource,
			TypeName: "aws_bedrock_model_invocation_logging_configuration",
			Name:     "Model Invocation Logging Configuration",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newProvisionedModelThroughputResource,
			TypeName: "aws_bedrock_provisioned_model_throughput",
			Name:     "Provisioned Model Throughput",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: "provisioned_model_arn",
			}),
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Bedrock
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*bedrock.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*bedrock.Options){
		bedrock.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *bedrock.Options) {
			if region := config[names.AttrRegion].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         p.ServicePackageName(),
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return bedrock.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*bedrock.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*bedrock.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *bedrock.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*bedrock.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

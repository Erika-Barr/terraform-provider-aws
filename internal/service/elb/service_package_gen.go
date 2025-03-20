// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package elb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceLoadBalancer,
			TypeName: "aws_elb",
			Name:     "Classic Load Balancer",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceHostedZoneID,
			TypeName: "aws_elb_hosted_zone_id",
			Name:     "Hosted Zone ID",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: false,
			},
		},
		{
			Factory:  dataSourceServiceAccount,
			TypeName: "aws_elb_service_account",
			Name:     "Service Account",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: false,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceAppCookieStickinessPolicy,
			TypeName: "aws_app_cookie_stickiness_policy",
			Name:     "App Cookie Stickiness Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceLoadBalancer,
			TypeName: "aws_elb",
			Name:     "Classic Load Balancer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceAttachment,
			TypeName: "aws_elb_attachment",
			Name:     "Attachment",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceCookieStickinessPolicy,
			TypeName: "aws_lb_cookie_stickiness_policy",
			Name:     "LB Cookie Stickiness Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceSSLNegotiationPolicy,
			TypeName: "aws_lb_ssl_negotiation_policy",
			Name:     "SSL Negotiation Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceBackendServerPolicy,
			TypeName: "aws_load_balancer_backend_server_policy",
			Name:     "Backend Server Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceListenerPolicy,
			TypeName: "aws_load_balancer_listener_policy",
			Name:     "Listener Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourcePolicy,
			TypeName: "aws_load_balancer_policy",
			Name:     "Load Balancer Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceProxyProtocolPolicy,
			TypeName: "aws_proxy_protocol_policy",
			Name:     "Proxy Protocol Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ELB
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*elasticloadbalancing.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*elasticloadbalancing.Options){
		elasticloadbalancing.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *elasticloadbalancing.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "elasticloadbalancing",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return elasticloadbalancing.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*elasticloadbalancing.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*elasticloadbalancing.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *elasticloadbalancing.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*elasticloadbalancing.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

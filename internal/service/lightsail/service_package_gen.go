// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{
		{
			Factory:  ResourceBucket,
			TypeName: "aws_lightsail_bucket",
			Name:     "Bucket",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Bucket",
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceBucketAccessKey,
			TypeName: "aws_lightsail_bucket_access_key",
			Name:     "Bucket Access Key",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceBucketResourceAccess,
			TypeName: "aws_lightsail_bucket_resource_access",
			Name:     "Bucket Resource Access",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceCertificate,
			TypeName: "aws_lightsail_certificate",
			Name:     "Certificate",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Certificate",
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceContainerService,
			TypeName: "aws_lightsail_container_service",
			Name:     "Container Service",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "ContainerService",
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceContainerServiceDeploymentVersion,
			TypeName: "aws_lightsail_container_service_deployment_version",
			Name:     "Container Service Deployment Version",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceDatabase,
			TypeName: "aws_lightsail_database",
			Name:     "Database",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Database",
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceDisk,
			TypeName: "aws_lightsail_disk",
			Name:     "Disk",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Disk",
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceDiskAttachment,
			TypeName: "aws_lightsail_disk_attachment",
			Name:     "Disk Attachment",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceDistribution,
			TypeName: "aws_lightsail_distribution",
			Name:     "Distribution",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Distribution",
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceDomain,
			TypeName: "aws_lightsail_domain",
			Name:     "Domain",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceDomainEntry,
			TypeName: "aws_lightsail_domain_entry",
			Name:     "Domain Entry",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_lightsail_instance",
			Name:     "Instance",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Instance",
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceInstancePublicPorts,
			TypeName: "aws_lightsail_instance_public_ports",
			Name:     "Instance Public Ports",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceKeyPair,
			TypeName: "aws_lightsail_key_pair",
			Name:     "KeyPair",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "KeyPair",
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceLoadBalancer,
			TypeName: "aws_lightsail_lb",
			Name:     "LB",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "LB",
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceLoadBalancerAttachment,
			TypeName: "aws_lightsail_lb_attachment",
			Name:     "Load Balancer Attachment",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceLoadBalancerCertificate,
			TypeName: "aws_lightsail_lb_certificate",
			Name:     "Load Balancer Certificate",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceLoadBalancerCertificateAttachment,
			TypeName: "aws_lightsail_lb_certificate_attachment",
			Name:     "Load Balancer Certificate Attachment",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceLoadBalancerHTTPSRedirectionPolicy,
			TypeName: "aws_lightsail_lb_https_redirection_policy",
			Name:     "Load Balancer HTTPS Redirection Policy",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceLoadBalancerStickinessPolicy,
			TypeName: "aws_lightsail_lb_stickiness_policy",
			Name:     "Load Balancer Stickiness Policy",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceStaticIP,
			TypeName: "aws_lightsail_static_ip",
			Name:     "Static IP",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceStaticIPAttachment,
			TypeName: "aws_lightsail_static_ip_attachment",
			Name:     "Static IP Attachment",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Lightsail
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*lightsail.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*lightsail.Options){
		lightsail.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *lightsail.Options) {
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

	return lightsail.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*lightsail.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*lightsail.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *lightsail.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*lightsail.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

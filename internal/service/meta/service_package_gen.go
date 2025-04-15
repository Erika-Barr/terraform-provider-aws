// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package meta

import (
	"context"
	"unique"

	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newARNDataSource,
			TypeName: "aws_arn",
			Name:     "ARN",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newDefaultTagsDataSource,
			TypeName: "aws_default_tags",
			Name:     "Default Tags",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newIPRangesDataSource,
			TypeName: "aws_ip_ranges",
			Name:     "IP Ranges",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newPartitionDataSource,
			TypeName: "aws_partition",
			Name:     "Partition",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newRegionDataSource,
			TypeName: "aws_region",
			Name:     "Region",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newRegionsDataSource,
			TypeName: "aws_regions",
			Name:     "Regions",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newServiceDataSource,
			TypeName: "aws_service",
			Name:     "Service",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newServicePrincipalDataSource,
			TypeName: "aws_service_principal",
			Name:     "Service Principal",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return "meta"
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

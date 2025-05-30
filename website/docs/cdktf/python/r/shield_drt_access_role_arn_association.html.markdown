---
subcategory: "Shield"
layout: "aws"
page_title: "AWS: aws_shield_drt_access_role_arn_association"
description: |-
  Terraform resource for managing an AWS Shield DRT Access Role ARN Association.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_shield_drt_access_role_arn_association

Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks.
For more information see [Configure AWS SRT Support](https://docs.aws.amazon.com/waf/latest/developerguide/authorize-srt.html)

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iam_role import IamRole
from imports.aws.iam_role_policy_attachment import IamRolePolicyAttachment
from imports.aws.shield_drt_access_role_arn_association import ShieldDrtAccessRoleArnAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = IamRole(self, "example",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "drt.shield.amazonaws.com"
                        },
                        "Sid": ""
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            name="example-role"
        )
        aws_iam_role_policy_attachment_example = IamRolePolicyAttachment(self, "example_1",
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy",
            role=example.name
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_iam_role_policy_attachment_example.override_logical_id("example")
        aws_shield_drt_access_role_arn_association_example =
        ShieldDrtAccessRoleArnAssociation(self, "example_2",
            role_arn=example.arn
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_shield_drt_access_role_arn_association_example.override_logical_id("example")
```

## Argument Reference

The following arguments are required:

* `role_arn` - (Required) The Amazon Resource Name (ARN) of the role the SRT will use to access your AWS account. Prior to making the AssociateDRTRole request, you must attach the `AWSShieldDRTAccessPolicy` managed policy to this role.

## Attribute Reference

This resource exports no additional attributes.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Shield DRT access role ARN association using the AWS account ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.shield_drt_access_role_arn_association import ShieldDrtAccessRoleArnAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ShieldDrtAccessRoleArnAssociation.generate_config_for_import(self, "example", "123456789012")
```

Using `terraform import`, import Shield DRT access role ARN association using the AWS account ID. For example:

```console
% terraform import aws_shield_drt_access_role_arn_association.example 123456789012
```

<!-- cache-key: cdktf-0.20.8 input-64153c99f6109f71d2f116c2ff3e7fcfe38fd4c4748e5390b47d3d064e3cbc4b -->
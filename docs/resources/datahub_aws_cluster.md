---
page_title: "cdp_datahub_aws_cluster Resource - terraform-provider-cdp"
subcategory: "datahub"
description: |-
  Creates an AWS Data hub cluster.
---

# cdp_datahub_aws_cluster (Resource)

Creates an AWS Data hub cluster.

## Example Usage

```terraform
// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

terraform {
  required_providers {
    cdp = {
      source = "registry.terraform.io/cloudera/cdp"
    }
  }
}

resource "cdp_datahub_aws_cluster" "aws-cluster" {
  name               = "<value>"
  environment        = "<value>"
  cluster_template   = "7.2.15 - Data Engineering: Apache Spark, Apache Hive, Apache Oozie"
  cluster_definition = "7.2.15 - Data Engineering for AWS"
  /* The below section kept here as a working example if one would like to use the cluster creation w/o the usage of the cluster definition

  /*instance_group = [
    {
      node_count = 0
      instance_group_name = "gateway"
      instance_group_type = "CORE"
      instance_type = "m5.2xlarge"
      root_volume_size = 100
      attached_volume_configuration = [
        {
          volume_size = 100
          volume_count = 1
          volume_type = "gp3"
        }
      ]
      recovery_mode = "MANUAL"
      volume_encryption = {
        encryption = false
      }
      recipes = []
    },
    {
      node_count = 1
      instance_group_name = "master"
      instance_group_type = "GATEWAY"
      instance_type = "m5.4xlarge"
      root_volume_size = 100
      attached_volume_configuration = [
        {
          volume_size = 100
          volume_count = 1
          volume_type = "gp3"
        }
      ]
      recovery_mode = "MANUAL"
      volume_encryption = {
        encryption = false
      }
      recipes = []
    },
    {
      node_count = 3
      instance_group_name = "worker"
      instance_group_type = "CORE"
      instance_type = "r5d.2xlarge"
      root_volume_size = 100
      attached_volume_configuration = [
        {
          volume_size = 300
          volume_count = 1
          volume_type = "gp3"
        }
      ]
      recovery_mode = "MANUAL"
      volume_encryption = {
        encryption = false
      }
      recipes = []
    },
    {
      node_count = 3
      instance_group_name = "compute"
      instance_group_type = "CORE"
      instance_type = "r5d.2xlarge"
      root_volume_size = 100
      attached_volume_configuration = [
        {
          volume_size = 300
          volume_count = 1
          volume_type = "ephemeral"
        }
      ]
      recovery_mode = "MANUAL"
      volume_encryption = {
        encryption = false
      }
      recipes = []
    }
  ]*/
}

output "cluster" {
  value = cdp_datahub_aws_cluster.aws-cluster.name
}

output "environment" {
  value = cdp_datahub_aws_cluster.aws-cluster.environment
}

output "cluster_template" {
  value = cdp_datahub_aws_cluster.aws-cluster.cluster_template
}

output "cluster_definition" {
  value = cdp_datahub_aws_cluster.aws-cluster.cluster_definition
}

/*
output "recipes" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].recipes
}

output "instance_group" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group
}

output "node_count" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].node_count
}

output "instance_group_name" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].instance_group_name
}

output "instance_group_type" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].instance_group_type
}

output "instance_type" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].instance_type
}

output "root_volume_size" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].root_volume_size
}

output "attached_volume_configuration" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].attached_volume_configuration
}

output "volume_size" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].attached_volume_configuration[*].volume_size
}

output "volume_count" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].attached_volume_configuration[*].volume_count
}

output "volume_type" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].attached_volume_configuration[*].volume_type
}

output "recovery_mode" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].recovery_mode
}

output "volume_encryption" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].volume_encryption
}

output "encryption" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].volume_encryption.encryption
}*/
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_definition` (String) The name of the cluster definition.
- `cluster_template` (String) The name of the cluster template.
- `environment` (String) The name of the environment where the cluster will belong to.
- `name` (String) The name of the cluster.

### Optional

- `instance_group` (Attributes List) (see [below for nested schema](#nestedatt--instance_group))

### Read-Only

- `crn` (String) The CRN of the cluster.
- `id` (String) The ID of this resource.

<a id="nestedatt--instance_group"></a>
### Nested Schema for `instance_group`

Required:

- `attached_volume_configuration` (Attributes List) Configuration regarding the attached volume to the specific instance group. (see [below for nested schema](#nestedatt--instance_group--attached_volume_configuration))
- `instance_group_name` (String) The name of the instance group.
- `instance_group_type` (String) The type of the instance group.
- `instance_type` (String) The cloud provider-side instance type.
- `node_count` (Number) The cluster node count. Has to be greater or equal than 0 and less than 100,000.
- `recovery_mode` (String) The type of the recovery mode.
- `root_volume_size` (Number) The size of the root volume in GB
- `volume_encryption` (Attributes) The volume encryption related configuration. (see [below for nested schema](#nestedatt--instance_group--volume_encryption))

Optional:

- `recipes` (Set of String) The set of recipe names that are going to be applied on the given instance group.

<a id="nestedatt--instance_group--attached_volume_configuration"></a>
### Nested Schema for `instance_group.attached_volume_configuration`

Required:

- `volume_count` (Number) The number of volumes to be attached.
- `volume_size` (Number) The size of the volume in GB.
- `volume_type` (String) The - cloud provider - type of the volume.


<a id="nestedatt--instance_group--volume_encryption"></a>
### Nested Schema for `instance_group.volume_encryption`

Required:

- `encryption` (Boolean)
---
layout: "vcd"
page_title: "Viettel IDC Tenant Manager: vcloud_tm_region_storage_policy"
sidebar_current: "docs-vcd-data-source-tm-region-storage-policy"
description: |-
  Provides a Viettel IDC Tenant Manager data source to read Region Storage Policies.
---

# vcloud\_tm\_region\_storage\_policy

Provides a Viettel IDC Tenant Manager data source to read Region Storage Policies.

This data source is exclusive to **Viettel IDC Tenant Manager**. Supported in provider *v4.0+*

-> To retrieve Storage Classes, use the [`vcloud_tm_storage_class`](/providers/viettelidc-provider/vcloud/latest/docs/data-sources/tm_storage_class)
data source instead

## Example Usage

```hcl
data "vcloud_tm_region" "region" {
  name = "my-region"
}

data "vcloud_tm_region_storage_policy" "sp" {
  region_id = data.vcloud_tm_region.region.id
  name      = "vSAN Default Storage Policy"
}

output "policy_id" {
  value = data.vcloud_tm_region_storage_policy.sp.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Region Storage Policy to read
* `region_id` - (Required) The ID of the Region where the Storage Policy belongs

## Attribute reference

* `description` - Description of the Region Storage Policy
* `status` - The creation status of the Region Storage Policy. Can be `NOT_READY` or `READY`
* `storage_capacity_mb` - Storage capacity in megabytes for this Region Storage Policy
* `storage_consumed_mb` - Consumed storage in megabytes for this Region Storage Policy
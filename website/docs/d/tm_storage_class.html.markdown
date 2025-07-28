---
layout: "vcd"
page_title: "Viettel IDC Tenant Manager: vcloud_tm_storage_class"
sidebar_current: "docs-vcd-data-source-tm-storage-class"
description: |-
  Provides a Viettel IDC Tenant Manager data source to read Storage Classes.
---

# vcloud\_tm\_storage\_class

Provides a Viettel IDC Tenant Manager data source to read Region Storage Classes.

This data source is exclusive to **Viettel IDC Tenant Manager**. Supported in provider *v4.0+*

## Example Usage

```hcl
data "vcloud_tm_region" "region" {
  name = "my-region"
}

data "vcloud_tm_region_storage_class" "sc" {
  region_id = data.vcloud_tm_region.region.id
  name      = "vSAN Default Storage Class"
}

resource "vcloud_tm_content_library" "cl" {
  name        = "My Library"
  description = "A simple library"
  storage_class_ids = [
    data.vcloud_tm_storage_class.sc.id
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Storage Class to read
* `region_id` - (Required) The ID of the Region where the Storage Class belongs

## Attribute reference

* `storage_capacity_mib` - The total storage capacity of the Storage Class in mebibytes
* `storage_consumed_mib` - For tenants, this represents the total storage given to all namespaces consuming from this
  Storage Class in mebibytes. For providers, this represents the total storage given to tenants from this Storage Class
  in mebibytes
* `zone_ids` - A set with all the IDs of the zones available to the Storage Class
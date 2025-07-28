---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_tier0_gateway"
sidebar_current: "docs-vcd-data-source-tm-tier0-gateway"
description: |-
  Provides a Viettel IDC Tenant Manager Tier 0 Gateway data source.
---

# vcloud\_tm\_tier0\_gateway

Provides a Viettel IDC Tenant Manager Tier 0 Gateway data source.

## Example Usage

```hcl
data "vcloud_tm_region" "demo" {
  name = "region-one"
}

data "vcloud_tm_tier0_gateway" "demo" {
  name      = "my-tier0-gateway"
  region_id = data.vcloud_tm_region.demo.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of TM Tier 0 Gateway originating in NSX-T 
* `region_id` - (Required) An ID of Region. Can be looked up using
  [vcloud_tm_region](/providers/viettelidc-provider/vcloud/latest/docs/data-sources/tm_region) data source

## Attribute Reference

* `description` - Description of the Tier 0 Gateway
* `parent_tier_0_id` - Parent Tier 0 Gateway ID if this is a Tier 0 VRF
* `already_imported` - Boolean flag if the Tier 0 Gateway is already consumed

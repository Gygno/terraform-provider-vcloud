---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_nsxt_tier0_router_interface"
sidebar_current: "docs-vcd-datasource-nsxt-tier0-router-interface"
description: |-
  Provides a data source to read NSX-T Tier-0 Router Interface that can be associated with IP Space Uplink
---

# vcloud\_nsxt\_tier0\_router\_interface

Supported in provider *v3.14+* with NSX-T IP Spaces.

Provides a data source to read NSX-T Tier-0 Router Interface that can be associated with IP Space
Uplink.

## Example Usage

```hcl
data "vcloud_nsxt_tier0_interface" "one" {
  external_network_id = vcloud_external_network_v2.provider-gateway.id
  name                = "interface-one"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of organization to which the edge gateway belongs. Optional if
  defined at provider level.
* `external_network_id` - (Required) The ID of Provider Gateway. Can be looked up using
  [vcloud_external_network_v2](/providers/viettelidc-provider/vcloud/latest/docs/data-sources/external_network_v2) data
  source.

## Attribute Reference

* `description` - The description of Tier-0 Router Interface in NSX-T.
* `type` - The type of Tier-0 Router Interface in NSX-T. One of `EXTERNAL`, `SERVICE` or `LOOPBACK`

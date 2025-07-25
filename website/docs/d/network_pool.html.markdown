---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_network_pool"
sidebar_current: "docs-vcd-data-source-network-pool"
description: |-
  Provides a data source for a network pool attached to a VCLOUD.
---

# vcloud\_network\_pool

Provides a data source for a network pool attached to a VCLOUD.

Supported in provider *v3.10+*

## Example Usage

```hcl
data "vcloud_network_pool" "np1" {
  name = "NSX-T Overlay 1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) network pool name.

## Attribute reference

All the attributes and arguments of the corresponding resource [vcloud_network_pool](/providers/viettelidc-provider/vcloud/latest/docs/resources/network_pool) are supported

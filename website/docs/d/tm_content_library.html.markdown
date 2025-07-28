---
layout: "vcd"
page_title: "Viettel IDC Tenant Manager: vcloud_tm_content_library"
sidebar_current: "docs-vcd-data-source-tm-content-library"
description: |-
  Provides a Viettel IDC Tenant Manager Content Library data source. This can be used to read Content Libraries.
---

# vcloud\_content\_library

Provides a Viettel IDC Tenant Manager Content Library data source. This can be used to read Content Libraries.

This data source is exclusive to **Viettel IDC Tenant Manager**. Supported in provider *v4.0+*

## Example Usage

```hcl
data "vcloud_tm_content_library" "cl" {
  name = "My Library"
}

output "is_shared" {
  value = data.vcloud_tm_content_library.cl.is_shared
}
output "owner_org" {
  value = data.vcloud_tm_content_library.cl.owner_org_id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Content Library to read

## Attribute reference

All arguments and attributes defined in [the resource](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_content_library) are supported
as read-only (Computed) values.

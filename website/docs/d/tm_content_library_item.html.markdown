---
layout: "vcd"
page_title: "Viettel IDC Tenant Manager: vcloud_tm_content_library_item"
sidebar_current: "docs-vcd-data-source-tm-content-library-item"
description: |-
  Provides a Viettel IDC Tenant Manager Content Library Item data source. This can be used to read Content Library Items.
---

# vcloud\_tm\_content\_library\_item

Provides a Viettel IDC Tenant Manager Content Library Item data source. This can be used to read Content Library Items.

This data source is exclusive to **Viettel IDC Tenant Manager**. Supported in provider *v4.0+*

## Example Usage

```hcl
data "vcloud_tm_content_library" "cl" {
  name = "My Library"
}

data "vcloud_tm_content_library_item" "cli" {
  name               = "My Library Item"
  content_library_id = data.vcloud_tm_content_library.cl.id
}

output "is_published" {
  value = data.vcloud_tm_content_library_item.cli.is_published
}
output "image_identifier" {
  value = data.vcloud_tm_content_library_item.cli.image_identifier
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Content Library to read
* `content_library_id` - (Required) ID of the Content Library that this item belongs to. Can be obtained with [a data source](/providers/viettelidc-provider/vcloud/latest/docs/data-sources/tm_content_library)

## Attribute reference

All arguments and attributes defined in [the resource](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_content_library_item) are supported
as read-only (Computed) values.

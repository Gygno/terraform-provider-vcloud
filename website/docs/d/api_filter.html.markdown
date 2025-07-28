---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_api_filter"
sidebar_current: "docs-vcd-data-source-api-filter"
description: |-
  Provides a data source to read API Filters. An API Filter allows to extend API with customised URLs
  that can be redirected to an External Endpoint.
---

# vcloud\_api\_filter

Supported in provider *v3.14+*.

Provides a data source to read API Filters. An API Filter allows to extend API with customised URLs
that can be redirected to an [`vcloud_external_endpoint`](/providers/viettelidc-provider/vcloud/latest/docs/resources/external_endpoint).

~> Only `System Administrator` can use this data source.

## Example Usage

```hcl
data "vcloud_api_filter" "api_filter1" {
  api_filter_id = "urn:vcloud:apiFilter:4252ab09-eed8-4bc6-86d7-6019090273f5"
}
```

## Argument Reference

The following arguments are supported:

* `api_filter_id` - (Required) ID of the API Filter. This is the only way of unequivocally identify an API Filter. A list of
available API Filters can be obtained by using the `list@` option of the import mechanism of the [resource](/providers/viettelidc-provider/vcloud/latest/docs/resources/api_filter#importing)

## Attribute Reference

All the arguments from [the `vcloud_api_filter` resource](/providers/viettelidc-provider/vcloud/latest/docs/resources/api_filter)
are available as read-only.

---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_nsxt_alb_virtual_service_http_resp_rules"
sidebar_current: "docs-vcd-datasource-nsxt-alb-virtual-service-http-resp-rules"
description: |-
  Provides a data source to read ALB Service Engine Groups policies for HTTP responses. HTTP response 
  rules can be used to to evaluate and modify the response and response attributes that the
  application returns.
---

# vcloud\_nsxt\_alb\_virtual\_service\_http\_resp\_rules

Supported in provider *v3.14+* with NSX-T and ALB.

Provides a data source to read ALB Service Engine Groups policies for HTTP responses. HTTP response 
rules can be used to to evaluate and modify the response and response attributes that the
application returns.

## Example Usage

```hcl
data "vcloud_nsxt_alb_virtual_service_http_resp_rules" "response-rules" {
  virtual_service_id = vcloud_nsxt_alb_virtual_service.test.id
}
```

## Argument Reference

The following arguments are supported:

* `virtual_service_id` - (Required) An ID of existing ALB Virtual Service.

## Attribute Reference

All the arguments and attributes defined in
[`vcloud_nsxt_alb_virtual_service_http_resp_rules`](/providers/viettelidc-provider/vcloud/latest/docs/resources/nsxt_alb_virtual_service_http_resp_rules)
resource are available.
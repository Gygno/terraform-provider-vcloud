---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_nsxt_alb_virtual_service_http_req_rules"
sidebar_current: "docs-vcd-datasource-nsxt-alb-virtual-service-http-req-rules"
description: |-
  Provides a data source to read ALB Service Engine Groups policies for HTTP requests. HTTP request 
  rules modify requests before they are either forwarded to the application, used as a basis for 
  content switching, or discarded.
---

# vcloud\_nsxt\_alb\_virtual\_service\_http\_req\_rules

Supported in provider *v3.14+* with NSX-T and ALB.

Provides a data source to read ALB Service Engine Groups policies for HTTP requests. HTTP request 
rules modify requests before they are either forwarded to the application, used as a basis for 
content switching, or discarded.

## Example Usage

```hcl
data "vcloud_nsxt_alb_virtual_service_http_req_rules" "request-rules" {
  virtual_service_id = vcloud_nsxt_alb_virtual_service.test.id
}
```

## Argument Reference

The following arguments are supported:

* `virtual_service_id` - (Required) An ID of existing ALB Virtual Service.

## Attribute Reference

All the arguments and attributes defined in
[`vcloud_nsxt_alb_virtual_service_http_req_rules`](/providers/viettelidc-provider/vcloud/latest/docs/resources/nsxt_alb_virtual_service_http_req_rules)
resource are available.
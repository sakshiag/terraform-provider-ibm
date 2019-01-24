---
layout: "ibm"
page_title: "IBM : lbaas"
sidebar_current: "docs-ibm-resource-lbaas"
description: |-
  Manages IBM Load balancer as a service.
---

# ibm\_lbaas

Provides a resource for a load balancer as a service. This allows a load balancer as a service to be created, updated, and deleted. Currently, only one subnet is supported.
 
## Example Usage

```hcl

resource "ibm_compute_vm_instance" "vm_instances" {
  count = "2"
  ....
}

resource "ibm_lbaas" "lbaas" {
  name        = "terraformLB"
  description = "delete this"
  subnets     = [1511875]

  protocols = [{
    frontend_protocol     = "HTTPS"
    frontend_port         = 443
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "round_robin"
    tls_certificate_id    = 11670
  },
    {
      frontend_protocol     = "HTTP"
      frontend_port         = 80
      backend_protocol      = "HTTP"
      backend_port          = 80
      load_balancing_method = "round_robin"
    },
  ]
}


```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The load balancer's name.
* `description` - (Optional, string) A description of the load balancer.
* `subnets` - (Required, array) The subnet where the load balancer will be provisioned. Only one subnet is supported.
* `protocols` - (Optional, array) A nested block describing the protocols assigned to load balancer. Nested `protocols` blocks have the following structure:
  * `frontend_protocol` - (Required, string) The frontend protocol. Accepted values are 'TCP', 'HTTP', and 'HTTPS'.
  * `frontend_port` - (Required, integer) The frontend protocol port number. The port number must be in the range of 1 - 65535.
  * `backend_protocol` - (Required, string) The backend protocol. Accepted values are 'TCP', 'HTTP', and 'HTTPS'.
  * `backend_port` - (Required, integer) The backend protocol port number. The port number must be in the range of 1 - 65535.
  * `load_balancing_method` - (Optional, string) The load balancing algorithm. Accepted values are 'round_robin', 'weighted_round_robin', and 'least_connection'. The default is 'round_robin'.
  * `session_stickiness` - (Optional, string) The SOURCE_IP for session stickiness.
  * `max_conn` - (Optional, integer) The maximum number of connections the listener can accept. The number must be 1 - 64000.
  * `tls_certificate_id` - (Optional, integer) The ID of the SSL/TLS certificate being used for a protocol. This ID should be specified when `frontend protocol` has a value of `HTTPS`.
* `wait_time_minutes` - (Optional, integer) The duration, expressed in minutes, to wait for the lbaas instance to become available before declaring it as created. It is also the same amount of time waited for deletion to finish. The default value is `90`.

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the created policy.
* `datacenter` - The datacenter where the load balancer is provisioned. This is based on the subnet chosen while creating load-balancer.
* `type` - Specifies whether a load balancer is `PUBLIC` or `PRIVATE`.
* `status` - Specifies the operation status of the load balancer as `ONLINE` or `OFFLINE`.
* `vip` - The virtual IP address of the load balancer.
* `protocol_id` - The UUID of a load balancer protocol.

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
)

func dataSourceIBMISVPC() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISVPCRead,

		Schema: map[string]*schema.Schema{
			isVPCDefaultNetworkACL: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCIsDefault: {
				Type:     schema.TypeBool,
				Computed: true,
			},

			isVPCName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isVPCResourceGroup: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCTags: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func dataSourceIBMISVPCRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	vpcC := network.NewVPCClient(sess)

	name := d.Get(isVPCName).(string)

	vpcs, _, err := vpcC.List(name)
	if err != nil {
		return err
	}

	for _, vpc := range vpcs {
		if vpc.Name == name {
			d.SetId(vpc.ID.String())
			d.Set("id", vpc.ID.String())
			d.Set(isVPCName, vpc.Name)
			d.Set(isVPCIsDefault, vpc.IsDefault)
			d.Set(isVPCStatus, vpc.Status)
			d.Set(isVPCTags, vpc.Tags)
			if vpc.DefaultNetworkACL != nil {
				d.Set(isVPCDefaultNetworkACL, vpc.DefaultNetworkACL.ID)
			} else {
				d.Set(isVPCDefaultNetworkACL, nil)
			}
			return nil
		}
	}
	return fmt.Errorf("No VPC found with name %s", name)
}

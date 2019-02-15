package ibm

import (
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/terraform/helper/schema"
	gouuid "github.com/satori/go.uuid"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

const (
	isSecurityGroupName          = "name"
	isSecurityGroupResourceGroup = "resource_group"
	isSecurityGroupVPC           = "vpc"
	isSecurityGroupRules         = "rules"
	isSecurityGroupTags          = "tags"
)

func resourceIBMISSecurityGroup() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMISSecurityGroupCreate,
		Read:     resourceIBMISSecurityGroupRead,
		Update:   resourceIBMISSecurityGroupUpdate,
		Delete:   resourceIBMISSecurityGroupDelete,
		Exists:   resourceIBMISSecurityGroupExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			isSecurityGroupName: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Security group name",
				//ValidateFunc: validateSecurityGroupId,
			},
			isSecurityGroupResourceGroup: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Security group's resource group id",
				ValidateFunc: validateResourceGroupId,
				ForceNew:     true,
			},
			isSecurityGroupVPC: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Security group's resource group id",
				ValidateFunc: validateVPCIdentity,
				ForceNew:     true,
			},

			isSecurityGroupRules: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Security Rules",
				Elem: &schema.Resource{
					Schema: makeIBMISSecurityRuleSchema(),
				},
			},

			isSecurityGroupTags: {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceIBMISSecurityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	parsed, err := parseIBMISSecurityGroupDictionary(d, "create")
	if err != nil {
		return err
	}
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)

	sgdef, err := makeIBMISSecurityGroupCreateParams(parsed)
	group, err := sgC.Create(sgdef)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	d.SetId(group.ID.String())
	return resourceIBMISSecurityGroupRead(d, meta)
}

func resourceIBMISSecurityGroupRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)

	group, err := sgC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set(isSecurityGroupName, group.Name)
	d.Set(isSecurityGroupVPC, group.Vpc.ID.String())
	if group.ResourceGroup != nil {
		d.Set(isSecurityGroupResourceGroup, group.ResourceGroup.ID.String())
	}
	if len(group.Tags) > 0 {
		tags := make([]string, 0)
		for _, t := range tags {
			tags = append(tags, t)
		}
		d.Set(isSecurityGroupTags, tags)
	}
	rules := make([]map[string]interface{}, 0)
	if len(group.Rules) > 0 {
		for _, rule := range group.Rules {
			r := make(map[string]interface{})
			if rule.Code != nil {
				r[isSecurityGroupRuleCode] = int(*rule.Code)
			}
			if rule.Type != nil {
				r[isSecurityGroupRuleType] = int(*rule.Type)
			}
			if rule.PortMin != nil {
				r[isSecurityGroupRulePortMin] = int(*rule.PortMin)
			}
			if rule.PortMax != nil {
				r[isSecurityGroupRulePortMax] = int(*rule.PortMax)
			}
			r[isSecurityGroupRuleDirection] = rule.Direction
			r[isSecurityGroupRuleIPVersion] = rule.IPVersion
			if rule.Protocol != nil {
				r[isSecurityGroupRuleProtocol] = *rule.Protocol
			}

			rules = append(rules, r)
		}
	}
	d.Set(isSecurityGroupRules, rules)
	d.SetId(group.ID.String())
	return nil
}

func resourceIBMISSecurityGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)
	if !d.HasChange(isSecurityGroupName) {
		return resourceIBMISSecurityGroupRead(d, meta)
	}

	name := d.Get(isSecurityGroupName).(string)

	_, err = sgC.Update(d.Id(), name)
	if err != nil {
		return err
	}
	err = resourceIBMISSecurityGroupRead(d, meta)
	return err
}

func resourceIBMISSecurityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)

	err = sgC.Delete(d.Id())
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return nil
			}
		}
		return err
	}
	return err
}

func resourceIBMISSecurityGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	sgC := network.NewSecurityGroupClient(sess)

	_, err = sgC.Get(d.Id())
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

type parsedIBMISSecurityGroupRule struct {
	// After parsing, unused string fields are set to
	// "" and unused int64 fields will be set to -1.
	// This ("" for unused strings and -1 for unused int64s)
	// is expected by our riaas API client.
	secgrpID       string
	ruleID         string
	direction      string
	ipversion      string
	remote         string
	remoteAddress  string
	remoteCIDR     string
	remoteSecGrpID string
	protocol       string
	icmpType       int64
	icmpCode       int64
	portMin        int64
	portMax        int64
}

func newParsedIBMISSecurityGroupRule() *parsedIBMISSecurityGroupRule {
	return &parsedIBMISSecurityGroupRule{
		icmpType: -1,
		icmpCode: -1,
		portMin:  -1,
		portMax:  -1,
	}
}

type parsedIBMISSecurityGroupDictionary struct {
	name          string
	resourceGroup string
	vpc           string
	tags          []string
}

func newParsedIBMISSecurityGroupDictionary() *parsedIBMISSecurityGroupDictionary {
	p := &parsedIBMISSecurityGroupDictionary{}
	p.tags = make([]string, 0)
	return p
}

func parseSecurityGroupTags(d *schema.ResourceData, p *parsedIBMISSecurityGroupDictionary) {
	tagSet := d.Get(isSecurityGroupTags).(*schema.Set)

	if tagSet.Len() == 0 {
		return
	}

	for _, elem := range tagSet.List() {
		p.tags = append(p.tags, elem.(string))
	}
	return
}

func parseIBMISSecurityGroupDictionary(d *schema.ResourceData, tag string) (*parsedIBMISSecurityGroupDictionary, error) {
	parsed := newParsedIBMISSecurityGroupDictionary()
	parsed.name = d.Get(isSecurityGroupName).(string)
	parsed.resourceGroup = d.Get(isSecurityGroupResourceGroup).(string)
	parsed.vpc = d.Get(isSecurityGroupVPC).(string)
	parseSecurityGroupTags(d, parsed)
	return parsed, nil
}

func makeStrfmtUUID(s string) (strfmt.UUID, error) {
	uuid, err := gouuid.FromString(s)
	if err != nil {
		return strfmt.UUID(""), err
	}
	return strfmt.UUID(uuid.String()), nil
}

func makeIBMISSecurityGroupCreateParams(parsed *parsedIBMISSecurityGroupDictionary) (*models.PostSecurityGroupsParamsBody, error) {
	params := &models.PostSecurityGroupsParamsBody{}
	params.Name = parsed.name
	if parsed.resourceGroup != "" {
		uuid, err := makeStrfmtUUID(parsed.resourceGroup)
		if err != nil {
			return nil, err
		}
		params.ResourceGroup = &models.PostSecurityGroupsParamsBodyResourceGroup{ID: uuid}
	}
	uuid, err := makeStrfmtUUID(parsed.vpc)
	if err != nil {
		return nil, err
	}
	params.Vpc = &models.PostSecurityGroupsParamsBodyVpc{ID: uuid}
	params.Tags = parsed.tags

	return params, nil
}

func makeIBMISSecurityRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{

		isSecurityGroupRuleDirection: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Direction of traffic to enforce, either ingress or egress",
		},

		isSecurityGroupRuleIPVersion: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "IP version: ipv4 or ipv6",
		},

		isSecurityGroupRuleRemote: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Security group id: an IP address, a CIDR block, or a single security group identifier",
		},

		isSecurityGroupRuleType: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRuleCode: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRulePortMin: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRulePortMax: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRuleProtocol: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

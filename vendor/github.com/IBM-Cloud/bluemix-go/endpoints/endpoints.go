package endpoints

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/helpers"
)

//EndpointLocator ...
type EndpointLocator interface {
	AccountManagementEndpoint() (string, error)
	CFAPIEndpoint() (string, error)
<<<<<<< HEAD
	MCCPAPIEndpoint() (string, error)
	ContainerEndpoint() (string, error)
	CisEndpoint() (string, error)
	IAMEndpoint() (string, error)
	IAMPAPEndpoint() (string, error)
=======
	ContainerEndpoint() (string, error)
	ContainerRegistryEndpoint() (string, error)
	CisEndpoint() (string, error)
	GlobalSearchEndpoint() (string, error)
	GlobalTaggingEndpoint() (string, error)
	IAMEndpoint() (string, error)
	IAMPAPEndpoint() (string, error)
	ICDEndpoint() (string, error)
	MCCPAPIEndpoint() (string, error)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	ResourceManagementEndpoint() (string, error)
	ResourceControllerEndpoint() (string, error)
	ResourceCatalogEndpoint() (string, error)
	UAAEndpoint() (string, error)
<<<<<<< HEAD
	ContainerRegistryEndpoint() (string, error)
=======
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}

const (
	//ErrCodeServiceEndpoint ...
	ErrCodeServiceEndpoint = "ServiceEndpointDoesnotExist"
)

var regionToEndpoint = map[string]map[string]string{
<<<<<<< HEAD
=======
	"account": {
		"us-south": "https://accountmanagement.ng.bluemix.net",
		"us-east":  "https://accountmanagement.us-east.bluemix.net",
		"eu-gb":    "https://accountmanagement.eu-gb.bluemix.net",
		"au-syd":   "https://accountmanagement.au-syd.bluemix.net",
		"eu-de":    "https://accountmanagement.eu-de.bluemix.net",
		"jp-tok":   "https://accountmanagement.jp-tok.bluemix.net",
	},
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	"cf": {
		"us-south": "https://api.ng.bluemix.net",
		"us-east":  "https://api.us-east.bluemix.net",
		"eu-gb":    "https://api.eu-gb.bluemix.net",
		"au-syd":   "https://api.au-syd.bluemix.net",
		"eu-de":    "https://api.eu-de.bluemix.net",
<<<<<<< HEAD
	},
	"mccp": {
		"us-south": "https://mccp.ng.bluemix.net",
		"us-east":  "https://mccp.us-east.bluemix.net",
		"eu-gb":    "https://mccp.eu-gb.bluemix.net",
		"au-syd":   "https://mccp.au-syd.bluemix.net",
		"eu-de":    "https://mccp.eu-de.bluemix.net",
	},
	"iam": {
		"us-south": "https://iam.bluemix.net",
		"us-east":  "https://iam.bluemix.net",
		"eu-gb":    "https://iam.bluemix.net",
		"au-syd":   "https://iam.bluemix.net",
		"eu-de":    "https://iam.bluemix.net",
		"jp-tok":   "https://iam.bluemix.net",
	},
	"iampap": {
		"us-south": "https://iam.bluemix.net",
		"us-east":  "https://iam.bluemix.net",
		"eu-gb":    "https://iam.bluemix.net",
		"au-syd":   "https://iam.bluemix.net",
		"eu-de":    "https://iam.bluemix.net",
		"jp-tok":   "https://iam.bluemix.net",
	},
	"uaa": {
		"us-south": "https://login.ng.bluemix.net/UAALoginServerWAR",
		"us-east":  "https://login.us-east.bluemix.net/UAALoginServerWAR",
		"eu-gb":    "https://login.eu-gb.bluemix.net/UAALoginServerWAR",
		"au-syd":   "https://login.au-syd.bluemix.net/UAALoginServerWAR",
		"eu-de":    "https://login.eu-de.bluemix.net/UAALoginServerWAR",
	},
	"account": {
		"us-south": "https://accountmanagement.ng.bluemix.net",
		"us-east":  "https://accountmanagement.us-east.bluemix.net",
		"eu-gb":    "https://accountmanagement.eu-gb.bluemix.net",
		"au-syd":   "https://accountmanagement.au-syd.bluemix.net",
		"eu-de":    "https://accountmanagement.eu-de.bluemix.net",
	},
	"cs": {
		"us-south": "https://containers.bluemix.net",
		"us-east":  "https://containers.bluemix.net",
		"eu-de":    "https://containers.bluemix.net",
		"au-syd":   "https://containers.bluemix.net",
		"eu-gb":    "https://containers.bluemix.net",
		"jp-tok":   "https://containers.bluemix.net",
=======
		"jp-tok":   "https://api.jp-tok.bluemix.net",
	},
	"cr": {
		"us-south": "https://registry.ng.bluemix.net",
		"us-east":  "https://registry.ng.bluemix.net",
		"eu-de":    "https://registry.eu-de.bluemix.net",
		"au-syd":   "https://registry.au-syd.bluemix.net",
		"eu-gb":    "https://registry.eu-gb.bluemix.net",
		"jp-tok":   "https://registry.jp-tok.bluemix.net",
	},
	"cs": {
		"us-south": "https://containers.cloud.ibm.com",
		"us-east":  "https://containers.cloud.ibm.com",
		"eu-de":    "https://containers.cloud.ibm.com",
		"au-syd":   "https://containers.cloud.ibm.com",
		"eu-gb":    "https://containers.cloud.ibm.com",
		"jp-tok":   "https://containers.cloud.ibm.com",
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	},
	"cis": {
		"us-south": "https://api.cis.cloud.ibm.com",
		"us-east":  "https://api.cis.cloud.ibm.com",
		"eu-de":    "https://api.cis.cloud.ibm.com",
		"au-syd":   "https://api.cis.cloud.ibm.com",
		"eu-gb":    "https://api.cis.cloud.ibm.com",
		"jp-tok":   "https://api.cis.cloud.ibm.com",
	},
<<<<<<< HEAD
	"resource-manager": {
		"us-south": "https://resource-manager.bluemix.net",
		"us-east":  "https://resource-manager.bluemix.net",
		"eu-de":    "https://resource-manager.bluemix.net",
		"au-syd":   "https://resource-manager.bluemix.net",
		"eu-gb":    "https://resource-manager.bluemix.net",
	},
	"resource-catalog": {
		"us-south": "https://resource-catalog.bluemix.net",
		"us-east":  "https://resource-catalog.bluemix.net",
		"eu-de":    "https://resource-catalog.bluemix.net",
		"au-syd":   "https://resource-catalog.bluemix.net",
		"eu-gb":    "https://resource-catalog.bluemix.net",
	},
	"resource-controller": {
		"us-south": "https://resource-controller.bluemix.net",
		"us-east":  "https://resource-controller.bluemix.net",
		"eu-de":    "https://resource-controller.bluemix.net",
		"au-syd":   "https://resource-controller.bluemix.net",
		"eu-gb":    "https://resource-controller.bluemix.net",
	},
	"cr": {
		"us-south": "https://registry.ng.bluemix.net",
		"us-east":  "https://registry.ng.bluemix.net",
		"eu-de":    "https://registry.eu-de.bluemix.net",
		"au-syd":   "https://registry.au-syd.bluemix.net",
		"eu-gb":    "https://registry.eu-gb.bluemix.net",
=======
	"global-search": {
		"us-south": "https://api.global-search-tagging.cloud.ibm.com",
		"us-east":  "https://api.global-search-tagging.cloud.ibm.com",
		"eu-de":    "https://api.global-search-tagging.cloud.ibm.com",
		"eu-gb":    "https://api.global-search-tagging.cloud.ibm.com",
		"au-syd":   "https://api.global-search-tagging.cloud.ibm.com",
		"jp-tok":   "https://api.global-search-tagging.cloud.ibm.com",
	},
	"global-tagging": {
		"us-south": "https://tags.global-search-tagging.cloud.ibm.com",
		"us-east":  "https://tags.global-search-tagging.cloud.ibm.com",
		"eu-de":    "https://tags.global-search-tagging.cloud.ibm.com",
		"eu-gb":    "https://tags.global-search-tagging.cloud.ibm.com",
		"au-syd":   "https://tags.global-search-tagging.cloud.ibm.com",
		"jp-tok":   "https://tags.global-search-tagging.cloud.ibm.com",
	},
	"iam": {
		"us-south": "https://iam.cloud.ibm.com",
		"us-east":  "https://iam.cloud.ibm.com",
		"eu-gb":    "https://iam.cloud.ibm.com",
		"au-syd":   "https://iam.cloud.ibm.com",
		"eu-de":    "https://iam.cloud.ibm.com",
		"jp-tok":   "https://iam.cloud.ibm.com",
	},
	"iampap": {
		"us-south": "https://iam.cloud.ibm.com",
		"us-east":  "https://iam.cloud.ibm.com",
		"eu-gb":    "https://iam.cloud.ibm.com",
		"au-syd":   "https://iam.cloud.ibm.com",
		"eu-de":    "https://iam.cloud.ibm.com",
		"jp-tok":   "https://iam.cloud.ibm.com",
	},
	"icd": {
		"us-south": "https://api.us-south.databases.cloud.ibm.com",
		"us-east":  "https://api.us-east.databases.cloud.ibm.com",
		"eu-de":    "https://api.eu-de.databases.cloud.ibm.com",
		"eu-gb":    "https://api.eu-gb.databases.cloud.ibm.com",
		"au-syd":   "https://api.au-syd.databases.cloud.ibm.com",
		"jp-tok":   "https://api.jp-tok.databases.cloud.ibm.com",
		"oslo01":   "https://api.osl01.databases.cloud.ibm.com",
	},
	"mccp": {
		"us-south": "https://mccp.ng.bluemix.net",
		"us-east":  "https://mccp.us-east.bluemix.net",
		"eu-gb":    "https://mccp.eu-gb.bluemix.net",
		"au-syd":   "https://mccp.au-syd.bluemix.net",
		"eu-de":    "https://mccp.eu-de.bluemix.net",
		"jp-tok":   "https://mccp.jp-tok.bluemix.net",
	},
	"resource-manager": {
		"us-south": "https://resource-controller.cloud.ibm.com",
		"us-east":  "https://resource-controller.cloud.ibm.com",
		"eu-de":    "https://resource-controller.cloud.ibm.com",
		"au-syd":   "https://resource-controller.cloud.ibm.com",
		"eu-gb":    "https://resource-controller.cloud.ibm.com",
		"jp-tok":   "https://resource-controller.cloud.ibm.com",
	},
	"resource-catalog": {
		"us-south": "https://globalcatalog.cloud.ibm.com",
		"us-east":  "https://globalcatalog.cloud.ibm.com",
		"eu-de":    "https://globalcatalog.cloud.ibm.com",
		"au-syd":   "https://globalcatalog.cloud.ibm.com",
		"eu-gb":    "https://globalcatalog.cloud.ibm.com",
		"jp-tok":   "https://globalcatalog.cloud.ibm.com",
	},
	"resource-controller": {
		"us-south": "https://resource-controller.cloud.ibm.com",
		"us-east":  "https://resource-controller.cloud.ibm.com",
		"eu-de":    "https://resource-controller.cloud.ibm.com",
		"au-syd":   "https://resource-controller.cloud.ibm.com",
		"eu-gb":    "https://resource-controller.cloud.ibm.com",
		"jp-tok":   "https://resource-controller.cloud.ibm.com",
	},
	"uaa": {
		"us-south": "https://login.ng.bluemix.net/UAALoginServerWAR",
		"us-east":  "https://login.us-east.bluemix.net/UAALoginServerWAR",
		"eu-gb":    "https://login.eu-gb.bluemix.net/UAALoginServerWAR",
		"au-syd":   "https://login.au-syd.bluemix.net/UAALoginServerWAR",
		"eu-de":    "https://login.eu-de.bluemix.net/UAALoginServerWAR",
		"jp-tok":   "https://login.jp-tok.bluemix.net/UAALoginServerWAR",
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	},
}

func init() {
	//TODO populate the endpoints which can be retrieved from given endpoints dynamically
	//Example - UAA can be found from the CF endpoint
}

type endpointLocator struct {
	region string
}

//NewEndpointLocator ...
func NewEndpointLocator(region string) EndpointLocator {
	return &endpointLocator{region: region}
}

<<<<<<< HEAD
=======
func (e *endpointLocator) AccountManagementEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["account"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_ACCOUNT_MANAGEMENT_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Account Management endpoint doesn't exist for region: %q", e.region))
}

>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
func (e *endpointLocator) CFAPIEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cf"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CF_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Cloud Foundry endpoint doesn't exist for region: %q", e.region))
}

<<<<<<< HEAD
func (e *endpointLocator) MCCPAPIEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["mccp"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_MCCP_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("MCCP API endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) UAAEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["uaa"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_UAA_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("UAA endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) AccountManagementEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["account"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_ACCOUNT_MANAGEMENT_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Account Management endpoint doesn't exist for region: %q", e.region))
=======
func (e *endpointLocator) ContainerEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cs"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Container Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) ContainerRegistryEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cr"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CR_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Container Registry Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) CisEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cis"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CIS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Cis Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) GlobalSearchEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["global-search"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_GS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Global Search Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) GlobalTaggingEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["global-tagging"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_GT_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Global Tagging Service endpoint doesn't exist for region: %q", e.region))
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}

func (e *endpointLocator) IAMEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["iam"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, ep), nil

	}
<<<<<<< HEAD
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("IAM  endpoint doesn't exist for region: %q", e.region))
=======
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("IAM endpoint doesn't exist for region: %q", e.region))
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}

func (e *endpointLocator) IAMPAPEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["iampap"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_IAMPAP_API_ENDPOINT"}, ep), nil

	}
<<<<<<< HEAD
	return "", fmt.Errorf("IAMPAP  endpoint doesn't exist for region: %q", e.region)
}

func (e *endpointLocator) ContainerEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cs"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Container Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) CisEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cis"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CIS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Cis Service endpoint doesn't exist for region: %q", e.region))
=======
	return "", fmt.Errorf("IAMPAP endpoint doesn't exist for region: %q", e.region)
}

func (e *endpointLocator) ICDEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["icd"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_ICD_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("ICD Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) MCCPAPIEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["mccp"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_MCCP_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("MCCP API endpoint doesn't exist for region: %q", e.region))
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}

func (e *endpointLocator) ResourceManagementEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["resource-manager"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_RESOURCE_MANAGEMENT_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Resource Management endpoint doesn't exist"))
}

func (e *endpointLocator) ResourceControllerEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["resource-controller"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_RESOURCE_CONTROLLER_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Resource Controller endpoint doesn't exist"))
}

func (e *endpointLocator) ResourceCatalogEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["resource-catalog"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_RESOURCE_CATALOG_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Resource Catalog endpoint doesn't exist"))
}

<<<<<<< HEAD
func (e *endpointLocator) ContainerRegistryEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cr"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CR_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Container Registry Service endpoint doesn't exist for region: %q", e.region))
=======
func (e *endpointLocator) UAAEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["uaa"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_UAA_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("UAA endpoint doesn't exist for region: %q", e.region))
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}

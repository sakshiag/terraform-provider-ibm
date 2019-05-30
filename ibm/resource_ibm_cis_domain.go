package ibm

import (
	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"reflect"
<<<<<<< HEAD
=======
	"strings"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
)

func resourceIBMCISDomain() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS object id",
				Required:    true,
			},
			"domain": {
				Type:        schema.TypeString,
				Description: "CISzone - Domain",
				Required:    true,
			},
			"paused": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"original_name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
<<<<<<< HEAD
		Create: resourceCISdomainCreate,
		Read:   resourceCISdomainRead,
		Update: resourceCISdomainUpdate,
		Delete: resourceCISdomainDelete,
		// No Exists due to errors in CIS API returning incorrect return codes not found (403)
=======
		Create:   resourceCISdomainCreate,
		Read:     resourceCISdomainRead,
		Update:   resourceCISdomainUpdate,
		Delete:   resourceCISdomainDelete,
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		Importer: &schema.ResourceImporter{},
	}
}

func resourceCISdomainCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
<<<<<<< HEAD
	log.Printf("   client %v\n", cisClient)
=======
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	if err != nil {
		return err
	}

	cisId := d.Get("cis_id").(string)
	zoneName := d.Get("domain").(string)

	zoneNew := v1.ZoneBody{Name: zoneName}
	var zoneObj v1.Zone
	var zone *v1.Zone
	zone, err = cisClient.Zones().CreateZone(cisId, zoneNew)
	if err != nil {
		log.Printf("CreateZones Failed %s\n", err)
		return err
	}
	zoneObj = *zone
	d.SetId(convertCisToTfTwoVar(zoneObj.Id, cisId))
	return resourceCISdomainRead(d, meta)
}

func resourceCISdomainRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
<<<<<<< HEAD

	zoneId, cisId, _ := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}

	log.Printf("resourceCISdomainRead - Getting Zone %v with cis_id %s\n", zoneId, cisId)
	var zone *v1.Zone

	zone, err = cisClient.Zones().GetZone(cisId, zoneId)
	if err != nil {
		log.Printf("resourceCISdomainRead - ListZones Failed %s\n", err)
=======
	zoneId, cisId, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	var zone *v1.Zone
	zone, err = cisClient.Zones().GetZone(cisId, zoneId)
	if err != nil {
		if checkCisZoneDeleted(d, meta, err, zone) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during DomainRead %v\n", err)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		return err
	}
	zoneObj := *zone
	d.Set("cis_id", cisId)
	d.Set("name", zoneObj.Name)
	d.Set("domain", zoneObj.Name)
	d.Set("status", zoneObj.Status)
	d.Set("paused", zoneObj.Paused)
	d.Set("name_servers", zoneObj.NameServers)
	d.Set("original_name_servers", zoneObj.OriginalNameServer)

	return nil
}
<<<<<<< HEAD
=======
func checkCisZoneDeleted(d *schema.ResourceData, meta interface{}, errCheck error, zone *v1.Zone) bool {
	// Check if error is due to removal of Cis resource and hence all subresources
	if strings.Contains(errCheck.Error(), "Object not found") ||
		strings.Contains(errCheck.Error(), "status code: 404") ||
		strings.Contains(errCheck.Error(), "status code: 400") ||
		strings.Contains(errCheck.Error(), "Invalid zone identifier") {
		log.Printf("[WARN] Removing resource from state because it's not found via the CIS API")
		return true
	}
	_, cisId, _ := convertTftoCisTwoVar(d.Id())
	exists, errNew := rcInstanceExists(cisId, "ibm_cis", meta)
	if errNew != nil {
		log.Printf("resourceCISdomainRead - Failure validating service exists %s\n", errNew)
		return false
	}
	if !exists {
		log.Printf("[WARN] Removing domain from state because parent cis instance is in removed state")
		return true
	}
	return false
}
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe

func resourceCISdomainUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCISdomainRead(d, meta)
}

func resourceCISdomainDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	zoneId, cisId, _ := convertTftoCisTwoVar(d.Id())
	var zone *v1.Zone
	emptyZone := new(v1.Zone)

<<<<<<< HEAD
	log.Println("Getting Zone to delete")
	zone, err = cisClient.Zones().GetZone(cisId, zoneId)
	if err != nil {
		log.Printf("GetZone Failed %s\n", err)
=======
	//Get Zone to delete
	zone, err = cisClient.Zones().GetZone(cisId, zoneId)
	if err != nil {
		if checkCisZoneDeleted(d, meta, err, zone) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during DomainRead %v\n", err)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		return err
	}

	zoneObj := *zone
	if !reflect.DeepEqual(emptyZone, zoneObj) {
<<<<<<< HEAD
		log.Println("Deleting Zone")
=======
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		err = cisClient.Zones().DeleteZone(cisId, zoneId)
		if err != nil {
			log.Printf("DeleteZone Failed %s\n", err)
			return err
		}
	}

	d.SetId("")
	return nil
}

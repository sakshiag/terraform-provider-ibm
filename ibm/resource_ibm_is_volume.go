package ibm

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/storage"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

const (
	isVolumeName           = "name"
	isVolumeResourceGroup  = "resource_group"
	isVolumeType           = "type"
	isVolumeZone           = "zone"
	isVolumeIOPs           = "iops"
	isVolumeCapacity       = "capacity"
	isVolumeProfile        = "profile"
	isVolumeSourceSnapshot = "source_snapshot"
	isVolumeStatus         = "status"
	isVolumeAutoDelete     = "auto_delete"
	isVolumeBillingTerm    = "billing_term"
	isVolumeTags           = "tags"
	isVolumeEncryptionKey  = "encryption_key"
)

func resourceIBMISVolume() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVolumeCreate,
		Read:     resourceIBMISVolumeRead,
		Update:   resourceIBMISVolumeUpdate,
		Delete:   resourceIBMISVolumeDelete,
		Exists:   resourceIBMISVolumeExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			isVolumeName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isVolumeType: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue([]string{"boot", "data"}),
			},

			isVolumeZone: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVolumeIOPs: {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},

			isVolumeCapacity: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
				ForceNew: false,
			},

			isVolumeProfile: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},

			isVolumeSourceSnapshot: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			isVolumeEncryptionKey: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			isVolumeAutoDelete: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},

			isVolumeBillingTerm: {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "hourly",
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue([]string{"hourly", "monthly"}),
			},

			isVolumeResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},

			isVolumeStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVolumeTags: {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: false,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceIBMISVolumeCreate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)

	name := d.Get(isVolumeName).(string)
	resgrp := d.Get(isVolumeResourceGroup).(string)
	capacity := d.Get(isVolumeCapacity).(int)
	iops := d.Get(isVolumeIOPs).(int)
	autodelete := d.Get(isVolumeAutoDelete).(bool)
	encryptionkey := d.Get(isVolumeEncryptionKey).(string)

	volType := strings.ToLower(d.Get(isVolumeType).(string))
	billingterm := models.PostServersParamsBodyBootVolumeAttachmentVolumeBillingTermHourly
	billingtermInput := strings.ToLower(d.Get(isVolumeBillingTerm).(string))
	if billingtermInput == "monthly" {
		billingterm = models.PostServersParamsBodyBootVolumeAttachmentVolumeBillingTermMonthly
	}
	profileName := d.Get(isVolumeProfile).(string)
	zone := d.Get(isVolumeZone).(string)
	snapshot := d.Get(isVolumeSourceSnapshot).(string)
	tags := d.Get(isVolumeTags).([]string)

	volume, err := volumeC.Create(name, capacity, iops, autodelete, encryptionkey,
		volType, billingterm, resgrp, tags, profileName, zone, snapshot)
	if err != nil {
		log.Printf("[DEBUG] Volume err %s", isErrorToString(err))
		return err
	}

	d.SetId(volume.ID.String())
	log.Printf("[INFO] Successfully create volume snapshot : %s", volume.ID.String())

	return resourceIBMISVolumeRead(d, meta)
}

func resourceIBMISVolumeRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)
	volumeid := d.Id()
	volume, err := volumeC.Get(volumeid)
	if err != nil {
		return err
	}

	d.SetId(volume.ID.String())
	d.Set(isVolumeName, volume.Name)
	if volume.ResourceGroup != nil {
		d.Set(isVolumeResourceGroup, volume.ResourceGroup)
	} else {
		d.Set(isVolumeResourceGroup, nil)
	}
	d.Set(isVolumeType, volume.Type)
	d.Set(isVolumeZone, volume.Zone.Name)
	d.Set(isVolumeIOPs, volume.Iops)
	d.Set(isVolumeCapacity, volume.Capacity)
	if volume.Profile != nil {
		d.Set(isVolumeProfile, volume.Profile.ID.String())
	}
	if volume.SourceSnapshot != nil {
		d.Set(isVolumeSourceSnapshot, volume.SourceSnapshot.ID.String())
	}
	d.Set(isVolumeStatus, volume.Status)
	d.Set(isVolumeAutoDelete, volume.AutoDelete)
	d.Set(isVolumeBillingTerm, volume.BillingTerm)
	d.Set(isVolumeTags, volume.Tags)

	return nil
}

func resourceIBMISVolumeUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)
	volumeid := d.Id()

	name := d.Get(isVolumeName).(string)
	capacity := d.Get(isVolumeCapacity).(int)
	iops := d.Get(isVolumeIOPs).(int)
	autodelete := d.Get(isVolumeAutoDelete).(bool)
	profileName := d.Get(isVolumeProfile).(string)
	_, err := volumeC.Update(volumeid, name, capacity, iops, &autodelete, profileName)
	if err != nil {
		log.Printf("[DEBUG] Volume err %s", isErrorToString(err))
		return err
	}
	return resourceIBMISVolumeRead(d, meta)
}

func resourceIBMISVolumeDelete(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)
	volumeid := d.Id()

	err := volumeC.Delete(volumeid)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMISVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)
	volumeid := d.Id()
	_, err := volumeC.Get(volumeid)
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

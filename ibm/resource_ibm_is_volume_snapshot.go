package ibm

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/storage"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isVolumeSnapshotVolumeID      = "volume_id"
	isVolumeSnapshotName          = "name"
	isVolumeSnapshotResourceGroup = "resource_group"
	isVolumeSnapshotStatus        = "status"
	isVolumeSnapshotTags          = "tags"
)

func resourceIBMISVolumeSnapshot() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVolumeSnapshotCreate,
		Read:     resourceIBMISVolumeSnapshotRead,
		Update:   resourceIBMISVolumeSnapshotUpdate,
		Delete:   resourceIBMISVolumeSnapshotDelete,
		Exists:   resourceIBMISVolumeSnapshotExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			isVolumeSnapshotName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isVolumeSnapshotVolumeID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVolumeSnapshotResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},

			isVolumeSnapshotStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVolumeSnapshotTags: {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceIBMISVolumeSnapshotCreate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)

	var resgrp string

	volumeid := d.Get(isVolumeSnapshotVolumeID).(string)
	name := d.Get(isVolumeSnapshotName).(string)
	if grp, ok := d.GetOk(isVolumeSnapshotResourceGroup); ok {
		resgrp = grp.(string)
	}

	snapshot, err := volumeC.CreateSnapshot(volumeid, name, resgrp, nil)
	if err != nil {
		log.Printf("[DEBUG] Volume Snapshot err %s", isErrorToString(err))
		return err
	}

	tid := makeTerraformRuleID(volumeid, snapshot.ID.String())

	d.SetId(tid)
	log.Printf("[INFO] Successfully create volume snapshot : %s", snapshot.ID.String())

	return resourceIBMISVolumeSnapshotRead(d, meta)
}

func resourceIBMISVolumeSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)

	volumeid, snapshotid, err := parseISTerraformID(d.Id())
	if err != nil {
		return err
	}

	snapshot, err := volumeC.GetSnapshot(volumeid, snapshotid)
	if err != nil {
		return err
	}

	d.Set(isVolumeSnapshotName, snapshot.Name)
	d.Set(isVolumeSnapshotStatus, snapshot.Status)
	if snapshot.ResourceGroup != nil {
		d.Set(isVolumeSnapshotResourceGroup, snapshot.ResourceGroup)
	} else {
		d.Set(isVolumeSnapshotResourceGroup, nil)
	}
	d.Set(isVolumeSnapshotTags, snapshot.Tags)

	return nil
}

func resourceIBMISVolumeSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)

	volumeid, snapshotid, err := parseISTerraformID(d.Id())
	if err != nil {
		return err
	}

	if d.HasChange(isVolumeSnapshotName) {
		name := d.Get(isVolumeSnapshotName).(string)
		_, err := volumeC.UpdateSnapshot(volumeid, snapshotid, name)
		if err != nil {
			return err
		}
	}

	return resourceIBMISVolumeSnapshotRead(d, meta)
}

func resourceIBMISVolumeSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)

	volumeid, snapshotid, err := parseISTerraformID(d.Id())
	if err != nil {
		return err
	}

	err = volumeC.DeleteSnapshot(volumeid, snapshotid)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMISVolumeSnapshotExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, _ := meta.(ClientSession).ISSession()
	volumeC := storage.NewVolumeClient(sess)

	volumeid, snapshotid, err := parseISTerraformID(d.Id())
	if err != nil {
		return false, err
	}

	_, err = volumeC.GetSnapshot(volumeid, snapshotid)
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

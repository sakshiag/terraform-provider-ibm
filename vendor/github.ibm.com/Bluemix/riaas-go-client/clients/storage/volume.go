package storage

import (
	"github.ibm.com/riaas/rias-api/riaas/client/storage"
	"github.ibm.com/riaas/rias-api/riaas/models"

	"github.com/go-openapi/strfmt"
	riaaserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
)

// VolumeClient ...
type VolumeClient struct {
	session *session.Session
}

// NewVolumeClient ...
func NewVolumeClient(sess *session.Session) *VolumeClient {
	return &VolumeClient{
		sess,
	}
}

// List ...
func (f *VolumeClient) List() (models.GetVolumesOKBodyVolumes, error) {
	return f.ListWithFilter("", "", "")
}

// ListWithFilter ...
func (f *VolumeClient) ListWithFilter(tag, zoneName, resourcegroupID string) (models.GetVolumesOKBodyVolumes, error) {
	params := storage.NewGetVolumesParams()
	if tag != "" {
		params = params.WithTag(&tag)
	}
	if zoneName != "" {
		params = params.WithZoneName(&zoneName)
	}

	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}

	resp, err := f.session.Riaas.Storage.GetVolumes(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload.Volumes, nil
}

// Get ...
func (f *VolumeClient) Get(id string) (*models.VolumeExtended, error) {
	params := storage.NewGetVolumesIDParams().WithID(id)
	resp, err := f.session.Riaas.Storage.GetVolumesID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *VolumeClient) Create(name string,
	capacity, iops int,
	autoDelete bool,
	encryptionKeyCRN string,
	volType, billingTerm, resourceGroupID string,
	tags []string, profileName, zoneName string,
	snapshotID string) (*models.Volume, error) {

	var body = models.PostVolumesParamsBody{
		Name:       name,
		Type:       volType,
		AutoDelete: &autoDelete,
	}

	if capacity != 0 {
		body.Capacity = int64(capacity)
	}

	if iops != 0 {
		body.Iops = int64(iops)
	}

	if encryptionKeyCRN != "" {
		body.EncryptionKey = &models.PostVolumesParamsBodyEncryptionKey{
			Crn: encryptionKeyCRN,
		}
	}

	if billingTerm != "" {
		body.BillingTerm = &billingTerm
	}

	if resourceGroupID != "" {
		body.ResourceGroup = &models.PostVolumesParamsBodyResourceGroup{
			ID: strfmt.UUID(resourceGroupID),
		}
	}

	if len(tags) != 0 {
		body.Tags = tags
	}

	if profileName != "" {
		body.Profile = &models.PostVolumesParamsBodyProfile{
			Name: profileName,
		}
	}

	if zoneName != "" {
		body.Zone = &models.NameReference{
			Name: zoneName,
		}
	}

	if snapshotID != "" {
		body.SourceSnapshot = &models.PostVolumesParamsBodySourceSnapshot{
			ID: strfmt.UUID(snapshotID),
		}
	}

	params := storage.NewPostVolumesParams().WithBody(&body)
	resp, err := f.session.Riaas.Storage.PostVolumes(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *VolumeClient) Delete(id string) error {
	params := storage.NewDeleteVolumesIDParams().WithID(id)
	_, err := f.session.Riaas.Storage.DeleteVolumesID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Update ...
func (f *VolumeClient) Update(id, name string, capacity, iops int, autoDelete *bool, profileName string) (*models.Volume, error) {
	var body = models.PatchVolumesIDParamsBody{}

	if name != "" {
		body.Name = name
	}

	if capacity != 0 {
		body.Capacity = int64(capacity)
	}

	if iops != 0 {
		body.Iops = int64(iops)
	}

	if autoDelete != nil {
		body.AutoDelete = *autoDelete
	}

	if profileName != "" {
		body.Profile = &models.PatchVolumesIDParamsBodyProfile{
			Name: profileName,
		}
	}

	params := storage.NewPatchVolumesIDParams().WithID(id).WithBody(&body)
	resp, err := f.session.Riaas.Storage.PatchVolumesID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// ListSnapshots ...
func (f *VolumeClient) ListSnapshots(volumeid string) (models.GetVolumesVolumeIDSnapshotsOKBodySnapshots, error) {
	return f.ListSnapshotsWithFilter(volumeid, "", "")
}

// ListSnapshotsWithFilter ...
func (f *VolumeClient) ListSnapshotsWithFilter(volumeid, resourcegroupID, tag string) (models.GetVolumesVolumeIDSnapshotsOKBodySnapshots, error) {
	params := storage.NewGetVolumesVolumeIDSnapshotsParams().WithVolumeID(volumeid)
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if tag != "" {
		params = params.WithTag(&tag)
	}
	resp, err := f.session.Riaas.Storage.GetVolumesVolumeIDSnapshots(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload.Snapshots, nil
}

// GetSnapshot ...
func (f *VolumeClient) GetSnapshot(volumeid, snapshotid string) (*models.VolumeSnapshot, error) {
	params := storage.NewGetVolumesVolumeIDSnapshotsIDParams().WithVolumeID(volumeid).WithID(snapshotid)
	resp, err := f.session.Riaas.Storage.GetVolumesVolumeIDSnapshotsID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteSnapshot ...
func (f *VolumeClient) DeleteSnapshot(volumeid, snapshotid string) error {
	params := storage.NewDeleteVolumesVolumeIDSnapshotsIDParams().WithVolumeID(volumeid).WithID(snapshotid)
	_, err := f.session.Riaas.Storage.DeleteVolumesVolumeIDSnapshotsID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// CreateSnapshot ...
func (f *VolumeClient) CreateSnapshot(volumeid, name, resourceGroupID string, tags []string) (*models.VolumeSnapshot, error) {

	var body = models.PostVolumesVolumeIDSnapshotsParamsBody{}
	if name != "" {
		body.Name = name
	}

	if resourceGroupID != "" {
		body.ResourceGroup = &models.PostVolumesVolumeIDSnapshotsParamsBodyResourceGroup{
			ID: strfmt.UUID(resourceGroupID),
		}
	}

	if len(tags) != 0 {
		body.Tags = tags
	}

	params := storage.NewPostVolumesVolumeIDSnapshotsParams().WithVolumeID(volumeid).WithBody(&body)
	resp, err := f.session.Riaas.Storage.PostVolumesVolumeIDSnapshots(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateSnapshot ...
func (f *VolumeClient) UpdateSnapshot(volumeid, snapshotid, newname string) (*models.VolumeSnapshot, error) {
	var body = models.PatchVolumesVolumeIDSnapshotsIDParamsBody{
		Name: newname,
	}

	params := storage.NewPatchVolumesVolumeIDSnapshotsIDParams().WithVolumeID(volumeid).WithID(snapshotid).WithBody(&body)
	resp, err := f.session.Riaas.Storage.PatchVolumesVolumeIDSnapshotsID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

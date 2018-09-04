package compute

import (
	"github.com/go-openapi/strfmt"

	"github.ibm.com/riaas/rias-api/riaas/client/compute"
	"github.ibm.com/riaas/rias-api/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// InstanceClient ...
type InstanceClient struct {
	session *session.Session
}

// NewInstanceClient ...
func NewInstanceClient(sess *session.Session) *InstanceClient {
	return &InstanceClient{
		sess,
	}
}

// ListWithFilter ...
func (f *InstanceClient) ListWithFilter(tag, zone, vpcid, subnetid, resourcegroupID, start string) (models.GetInstancesOKBodyInstances, string, error) {
	params := compute.NewGetInstancesParams()

	if tag != "" {
		params = params.WithTag(&tag)
	}
	if zone != "" {
		params = params.WithZoneName(&zone)
	}
	if vpcid != "" {
		params = params.WithVpcID(&vpcid)
	}
	if subnetid != "" {
		params = params.WithNetworkInterfacesSubnetID(&subnetid)
	}
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if start != "" {
		params = params.WithStart(&start)
	}

	resp, err := f.session.Riaas.Compute.GetInstances(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	if resp.Payload.Instances != nil {
		return resp.Payload.Instances, utils.GetNext(resp.Payload.Next), nil
	}
	return models.GetInstancesOKBodyInstances(resp.Payload.Servers), utils.GetNext(resp.Payload.Next), nil
}

// List ...
func (f *InstanceClient) List(start string) (models.GetInstancesOKBodyInstances, string, error) {
	return f.ListWithFilter("", "", "", "", "", start)
}

// Get ...
func (f *InstanceClient) Get(id string) (*models.Instance, error) {
	params := compute.NewGetInstancesIDParams().WithID(id)
	resp, err := f.session.Riaas.Compute.GetInstancesID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetInitParms ...
func (f *InstanceClient) GetInitParms(id string) (*models.InstanceInitialization, error) {
	params := compute.NewGetInstancesInstanceIDInitializationParams().WithInstanceID(id)
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDInitialization(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *InstanceClient) Create(instancedef *models.PostInstancesParamsBody) (*models.Instance, error) {
	params := compute.NewPostInstancesParams().WithBody(instancedef)
	resp, err := f.session.Riaas.Compute.PostInstances(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// CreateEasy ...
func (f *InstanceClient) CreateEasy(name string) (*models.Instance, error) {

	var body = models.PostInstancesParamsBody{
		Name: name,
	}
	return f.Create(&body)
}

// Update ...
func (f *InstanceClient) Update(id, name, profileName string) (*models.Instance, error) {
	var body = models.PatchInstancesIDParamsBody{}
	if name != "" {
		body.Name = name
	}
	if profileName != "" {
		var profile = models.PatchInstancesIDParamsBodyProfile{
			Name: profileName,
		}
		body.Profile = &profile
	}

	params := compute.NewPatchInstancesIDParams().WithID(id).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PatchInstancesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *InstanceClient) Delete(id string) error {
	params := compute.NewDeleteInstancesIDParams().WithID(id)
	_, err := f.session.Riaas.Compute.DeleteInstancesID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// CreateAction ...
func (f *InstanceClient) CreateAction(instanceid, actiontype string) (*models.InstanceAction, error) {
	body := models.PostInstancesInstanceIDActionsParamsBody{
		Type: actiontype,
	}
	params := compute.NewPostInstancesInstanceIDActionsParams().WithInstanceID(instanceid).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PostInstancesInstanceIDActions(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// ListActions ...
func (f *InstanceClient) ListActions(instanceid, start string) (models.GetInstancesInstanceIDActionsOKBodyActions, string, error) {
	params := compute.NewGetInstancesInstanceIDActionsParams().WithInstanceID(instanceid)
	if start != "" {
		params = params.WithStart(&start)
	}
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDActions(params, session.Auth(f.session))
	if err != nil {
		return nil, "", errors.ToError(err)
	}
	return resp.Payload.Actions, utils.GetNext(resp.Payload.Next), nil
}

// GetAction ...
func (f *InstanceClient) GetAction(instanceid, actionid string) (*models.InstanceAction, error) {
	params := compute.NewGetInstancesInstanceIDActionsIDParams().WithInstanceID(instanceid).WithID(actionid)
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDActionsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteAction ...
func (f *InstanceClient) DeleteAction(instanceid, actionid string) error {
	params := compute.NewDeleteInstancesInstanceIDActionsIDParams().WithInstanceID(instanceid).WithID(actionid)
	_, err := f.session.Riaas.Compute.DeleteInstancesInstanceIDActionsID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// ListProfiles ...
func (f *InstanceClient) ListProfiles(start string) (models.GetInstanceProfilesOKBodyProfiles, string, error) {
	params := compute.NewGetInstanceProfilesParams()
	resp, err := f.session.Riaas.Compute.GetInstanceProfiles(params, session.Auth(f.session))
	if err != nil {
		return nil, "", errors.ToError(err)
	}
	return resp.Payload.Profiles, utils.GetNext(resp.Payload.Next), nil
}

// GetProfile ...
func (f *InstanceClient) GetProfile(profileName string) (*models.InstanceProfile, error) {
	params := compute.NewGetInstanceProfilesNameParams().WithName(profileName)
	resp, err := f.session.Riaas.Compute.GetInstanceProfilesName(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// ListInterfacesWithFilter ...
func (f *InstanceClient) ListInterfacesWithFilter(instanceid, resourcegroupID, tag string) (models.GetInstancesInstanceIDNetworkInterfacesOKBodyNetworkInterfaces, error) {
	params := compute.NewGetInstancesInstanceIDNetworkInterfacesParams().WithInstanceID(instanceid)
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if tag != "" {
		params = params.WithTag(&tag)
	}

	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDNetworkInterfaces(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.NetworkInterfaces, nil
}

// ListInterfaces ...
func (f *InstanceClient) ListInterfaces(instanceid string) (models.GetInstancesInstanceIDNetworkInterfacesOKBodyNetworkInterfaces, error) {
	return f.ListInterfacesWithFilter(instanceid, "", "")
}

// GetInterface ...
func (f *InstanceClient) GetInterface(instanceid, interfaceid string) (*models.InstanceNetworkInterface, error) {
	params := compute.NewGetInstancesInstanceIDNetworkInterfacesIDParams().
		WithInstanceID(instanceid).
		WithID(interfaceid)
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// AddInterface ...
func (f *InstanceClient) AddInterface(instanceid, name, subnetID string, portSpeed int, v4address, v6address string,
	secondaryAddresses, securityGroupIDs []string, resourcegroupID string, tags []string) (*models.InstanceNetworkInterface, error) {

	body := models.PostInstancesInstanceIDNetworkInterfacesParamsBody{}
	body.Name = name
	body.PortSpeed = int64(portSpeed)
	if v6address != "" {
		body.PrimaryIPV6Address = v6address
	}
	if v4address != "" {
		body.PrimaryIPV4Address = v4address
	}

	if len(secondaryAddresses) != 0 {
		body.SecondaryAddresses = secondaryAddresses
	}

	if len(securityGroupIDs) != 0 {
		sgs := make([]*models.PostInstancesInstanceIDNetworkInterfacesParamsBodySecurityGroupsItems, len(securityGroupIDs))
		for i, sgid := range securityGroupIDs {
			sgref := models.PostInstancesInstanceIDNetworkInterfacesParamsBodySecurityGroupsItems{
				ID: strfmt.UUID(sgid),
			}
			sgs[i] = &sgref
		}
		body.SecurityGroups = sgs
	}

	subnetref := models.PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet{
		ID: strfmt.UUID(subnetID),
	}
	body.Subnet = &subnetref

	if resourcegroupID != "" {
		rgref := models.PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup{
			ID: strfmt.UUID(resourcegroupID),
		}
		body.ResourceGroup = &rgref
	}

	if len(tags) != 0 {
		body.Tags = tags
	}

	params := compute.NewPostInstancesInstanceIDNetworkInterfacesParams().WithInstanceID(instanceid).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PostInstancesInstanceIDNetworkInterfaces(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteInterface ...
func (f *InstanceClient) DeleteInterface(instanceid, interfaceid string) error {
	params := compute.NewDeleteInstancesInstanceIDNetworkInterfacesIDParams().WithInstanceID(instanceid).WithID(interfaceid)
	_, err := f.session.Riaas.Compute.DeleteInstancesInstanceIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// UpdateInterface ...
func (f *InstanceClient) UpdateInterface(instanceid, interfaceid, name string, portSpeed int) (*models.InstanceNetworkInterface, error) {

	body := models.PatchInstancesInstanceIDNetworkInterfacesIDParamsBody{}
	if name != "" {
		body.Name = name
	}
	if portSpeed != 0 {
		body.PortSpeed = int64(portSpeed)
	}
	params := compute.NewPatchInstancesInstanceIDNetworkInterfacesIDParams().WithInstanceID(instanceid).WithID(interfaceid).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PatchInstancesInstanceIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// ListInterfaceFloatingIPs ...
func (f *InstanceClient) ListInterfaceFloatingIPs(instanceid, interfaceid string) (models.GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBodyFloatingIps, error) {
	params := compute.NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsParams().
		WithInstanceID(instanceid).WithNetworkInterfaceID(interfaceid)
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIps(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.FloatingIps, nil
}

// GetInterfaceFloatingIP ...
func (f *InstanceClient) GetInterfaceFloatingIP(instanceid, interfaceid, address string) (*models.FloatingIP, error) {
	params := compute.NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams().
		WithInstanceID(instanceid).WithNetworkInterfaceID(interfaceid).WithAddress(address)
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// AddInterfaceFloatingIP ...
func (f *InstanceClient) AddInterfaceFloatingIP(instanceid, interfaceid, address string) (*models.FloatingIP, error) {
	params := compute.NewPutInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams().
		WithInstanceID(instanceid).WithNetworkInterfaceID(interfaceid).WithAddress(address)
	resp, err := f.session.Riaas.Compute.PutInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// RemoveInterfaceFloatingIP ...
func (f *InstanceClient) RemoveInterfaceFloatingIP(instanceid, interfaceid, address string) error {
	params := compute.NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams().
		WithInstanceID(instanceid).WithNetworkInterfaceID(interfaceid).WithAddress(address)
	_, err := f.session.Riaas.Compute.DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress(
		params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// ListVolAttachmentsWithFilter ...
func (f *InstanceClient) ListVolAttachmentsWithFilter(instanceid, resourcegroupID, tag string) (models.GetInstancesInstanceIDVolumeAttachmentsOKBodyVolumeAttachments, error) {
	params := compute.NewGetInstancesInstanceIDVolumeAttachmentsParams().WithInstanceID(instanceid)
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if tag != "" {
		params = params.WithTag(&tag)
	}

	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDVolumeAttachments(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.VolumeAttachments, nil
}

// ListVolAttachments ...
func (f *InstanceClient) ListVolAttachments(instanceid string) (models.GetInstancesInstanceIDVolumeAttachmentsOKBodyVolumeAttachments, error) {
	return f.ListVolAttachmentsWithFilter(instanceid, "", "")
}

// GetVolAttachment ...
func (f *InstanceClient) GetVolAttachment(instanceid, volAttachID string) (*models.InstanceVolumeAttachment, error) {
	params := compute.NewGetInstancesInstanceIDVolumeAttachmentsIDParams().
		WithInstanceID(instanceid).
		WithID(volAttachID)
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDVolumeAttachmentsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// AttachVolume ...
func (f *InstanceClient) AttachVolume(instanceid, volumeID, name string, resourcegroupID string, tags []string) (*models.InstanceVolumeAttachment, error) {

	body := models.PostInstancesInstanceIDVolumeAttachmentsParamsBody{}
	if name != "" {
		body.Name = name
	}
	body.Volume = &models.PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume{
		ID: strfmt.UUID(volumeID),
	}

	if resourcegroupID != "" {
		rgref := models.PostInstancesInstanceIDVolumeAttachmentsParamsBodyResourceGroup{
			ID: strfmt.UUID(resourcegroupID),
		}
		body.ResourceGroup = &rgref
	}

	if len(tags) != 0 {
		body.Tags = tags
	}

	params := compute.NewPostInstancesInstanceIDVolumeAttachmentsParams().WithInstanceID(instanceid).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PostInstancesInstanceIDVolumeAttachments(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteVolAttachment ...
func (f *InstanceClient) DeleteVolAttachment(instanceid, volAttachID string) error {
	params := compute.NewDeleteInstancesInstanceIDVolumeAttachmentsIDParams().WithInstanceID(instanceid).WithID(volAttachID)
	_, err := f.session.Riaas.Compute.DeleteInstancesInstanceIDVolumeAttachmentsID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// UpdateVolAttachment ...
func (f *InstanceClient) UpdateVolAttachment(instanceid, volAttachID, name string) (*models.InstanceVolumeAttachment, error) {

	body := models.PatchInstancesInstanceIDVolumeAttachmentsIDParamsBody{}
	if name != "" {
		body.Name = name
	}

	params := compute.NewPatchInstancesInstanceIDVolumeAttachmentsIDParams().WithInstanceID(instanceid).WithID(volAttachID).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PatchInstancesInstanceIDVolumeAttachmentsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

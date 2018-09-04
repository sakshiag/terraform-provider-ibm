package compute

import (
	"github.com/go-openapi/strfmt"

	"github.ibm.com/riaas/rias-api/riaas/client/compute"
	"github.ibm.com/riaas/rias-api/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
)

// ServerClient ...
type ServerClient struct {
	session *session.Session
}

// NewServerClient ...
func NewServerClient(sess *session.Session) *ServerClient {
	return &ServerClient{
		sess,
	}
}

// ListWithFilter ...
func (f *ServerClient) ListWithFilter(tag, zone, vpcid, subnetid, resourcegroupID string) (models.GetServersOKBodyServers, error) {
	params := compute.NewGetServersParams()

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

	resp, err := f.session.Riaas.Compute.GetServers(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload.Servers, nil
}

// List ...
func (f *ServerClient) List() (models.GetServersOKBodyServers, error) {
	return f.ListWithFilter("", "", "", "", "")
}

// Get ...
func (f *ServerClient) Get(id string) (*models.Server, error) {
	params := compute.NewGetServersIDParams().WithID(id)
	resp, err := f.session.Riaas.Compute.GetServersID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *ServerClient) Create(serverdef *models.PostServersParamsBody) (*models.Server, error) {
	params := compute.NewPostServersParams().WithBody(serverdef)
	resp, err := f.session.Riaas.Compute.PostServers(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// CreateEasy ...
func (f *ServerClient) CreateEasy(name string) (*models.Server, error) {

	var body = models.PostServersParamsBody{
		Name: name,
	}
	return f.Create(&body)
}

// Update ...
func (f *ServerClient) Update(id, name, profileName string) (*models.Server, error) {
	var body = models.PatchServersIDParamsBody{}
	if name != "" {
		body.Name = name
	}
	if profileName != "" {
		var profile = models.PatchServersIDParamsBodyProfile{
			Name: profileName,
		}
		body.Profile = &profile
	}

	params := compute.NewPatchServersIDParams().WithID(id).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PatchServersID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *ServerClient) Delete(id string) error {
	params := compute.NewDeleteServersIDParams().WithID(id)
	_, err := f.session.Riaas.Compute.DeleteServersID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// CreateAction ...
func (f *ServerClient) CreateAction(serverid, actiontype string) (*models.ServerAction, error) {
	body := models.PostServersServerIDActionsParamsBody{
		Type: actiontype,
	}
	params := compute.NewPostServersServerIDActionsParams().WithServerID(serverid).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PostServersServerIDActions(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// ListActions ...
func (f *ServerClient) ListActions(serverid string) (models.GetServersServerIDActionsOKBodyActions, error) {
	params := compute.NewGetServersServerIDActionsParams().WithServerID(serverid)
	resp, err := f.session.Riaas.Compute.GetServersServerIDActions(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.Actions, nil
}

// GetAction ...
func (f *ServerClient) GetAction(serverid, actionid string) (*models.ServerAction, error) {
	params := compute.NewGetServersServerIDActionsIDParams().WithServerID(serverid).WithID(actionid)
	resp, err := f.session.Riaas.Compute.GetServersServerIDActionsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteAction ...
func (f *ServerClient) DeleteAction(serverid, actionid string) error {
	params := compute.NewDeleteServersServerIDActionsIDParams().WithServerID(serverid).WithID(actionid)
	_, err := f.session.Riaas.Compute.DeleteServersServerIDActionsID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// ListProfiles ...
func (f *ServerClient) ListProfiles() (models.GetServerProfilesOKBodyProfiles, error) {
	params := compute.NewGetServerProfilesParams()
	resp, err := f.session.Riaas.Compute.GetServerProfiles(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.Profiles, nil
}

// GetProfile ...
func (f *ServerClient) GetProfile(profileName string) (*models.ServerProfile, error) {
	params := compute.NewGetServerProfilesNameParams().WithName(profileName)
	resp, err := f.session.Riaas.Compute.GetServerProfilesName(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// ListInterfacesWithFilter ...
func (f *ServerClient) ListInterfacesWithFilter(serverid, resourcegroupID, tag string) (models.GetServersServerIDNetworkInterfacesOKBodyNetworkInterfaces, error) {
	params := compute.NewGetServersServerIDNetworkInterfacesParams().WithServerID(serverid)
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if tag != "" {
		params = params.WithTag(&tag)
	}

	resp, err := f.session.Riaas.Compute.GetServersServerIDNetworkInterfaces(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.NetworkInterfaces, nil
}

// ListInterfaces ...
func (f *ServerClient) ListInterfaces(serverid string) (models.GetServersServerIDNetworkInterfacesOKBodyNetworkInterfaces, error) {
	return f.ListInterfacesWithFilter(serverid, "", "")
}

// GetInterface ...
func (f *ServerClient) GetInterface(serverid, interfaceid string) (*models.ServerNetworkInterface, error) {
	params := compute.NewGetServersServerIDNetworkInterfacesIDParams().
		WithServerID(serverid).
		WithID(interfaceid)
	resp, err := f.session.Riaas.Compute.GetServersServerIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// AddInterface ...
func (f *ServerClient) AddInterface(serverid, name, subnetID string, portSpeed int, v4address, v6address string,
	secondaryAddresses, securityGroupIDs []string, resourcegroupID string, tags []string) (*models.ServerNetworkInterface, error) {

	body := models.PostServersServerIDNetworkInterfacesParamsBody{}
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
		sgs := make([]*models.PostServersServerIDNetworkInterfacesParamsBodySecurityGroupsItems, len(securityGroupIDs))
		for i, sgid := range securityGroupIDs {
			sgref := models.PostServersServerIDNetworkInterfacesParamsBodySecurityGroupsItems{
				ID: strfmt.UUID(sgid),
			}
			sgs[i] = &sgref
		}
		body.SecurityGroups = sgs
	}

	subnetref := models.PostServersServerIDNetworkInterfacesParamsBodySubnet{
		ID: strfmt.UUID(subnetID),
	}
	body.Subnet = &subnetref

	if resourcegroupID != "" {
		rgref := models.PostServersServerIDNetworkInterfacesParamsBodyResourceGroup{
			ID: strfmt.UUID(resourcegroupID),
		}
		body.ResourceGroup = &rgref
	}

	if len(tags) != 0 {
		body.Tags = tags
	}

	params := compute.NewPostServersServerIDNetworkInterfacesParams().WithServerID(serverid).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PostServersServerIDNetworkInterfaces(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteInterface ...
func (f *ServerClient) DeleteInterface(serverid, interfaceid string) error {
	params := compute.NewDeleteServersServerIDNetworkInterfacesIDParams().WithServerID(serverid).WithID(interfaceid)
	_, err := f.session.Riaas.Compute.DeleteServersServerIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// UpdateInterface ...
func (f *ServerClient) UpdateInterface(serverid, interfaceid, name string, portSpeed int) (*models.ServerNetworkInterface, error) {

	body := models.PatchServersServerIDNetworkInterfacesIDParamsBody{}
	if name != "" {
		body.Name = name
	}
	if portSpeed != 0 {
		body.PortSpeed = int64(portSpeed)
	}
	params := compute.NewPatchServersServerIDNetworkInterfacesIDParams().WithServerID(serverid).WithID(interfaceid).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PatchServersServerIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// ListInterfaceFloatingIPs ...
func (f *ServerClient) ListInterfaceFloatingIPs(serverid, interfaceid string) (models.GetServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBodyFloatingIps, error) {
	params := compute.NewGetServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsParams().
		WithServerID(serverid).WithNetworkInterfaceID(interfaceid)
	resp, err := f.session.Riaas.Compute.GetServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIps(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.FloatingIps, nil
}

// GetInterfaceFloatingIP ...
func (f *ServerClient) GetInterfaceFloatingIP(serverid, interfaceid, address string) (*models.FloatingIP, error) {
	params := compute.NewGetServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams().
		WithServerID(serverid).WithNetworkInterfaceID(interfaceid).WithAddress(address)
	resp, err := f.session.Riaas.Compute.GetServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// AddInterfaceFloatingIP ...
func (f *ServerClient) AddInterfaceFloatingIP(serverid, interfaceid, address string) (*models.FloatingIP, error) {
	params := compute.NewPutServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams().
		WithServerID(serverid).WithNetworkInterfaceID(interfaceid).WithAddress(address)
	resp, err := f.session.Riaas.Compute.PutServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// RemoveInterfaceFloatingIP ...
func (f *ServerClient) RemoveInterfaceFloatingIP(serverid, interfaceid, address string) error {
	params := compute.NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams().
		WithServerID(serverid).WithNetworkInterfaceID(interfaceid).WithAddress(address)
	_, err := f.session.Riaas.Compute.DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress(
		params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// ListVolAttachmentsWithFilter ...
func (f *ServerClient) ListVolAttachmentsWithFilter(serverid, resourcegroupID, tag string) (models.GetServersServerIDVolumeAttachmentsOKBodyVolumeAttachments, error) {
	params := compute.NewGetServersServerIDVolumeAttachmentsParams().WithServerID(serverid)
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if tag != "" {
		params = params.WithTag(&tag)
	}

	resp, err := f.session.Riaas.Compute.GetServersServerIDVolumeAttachments(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.VolumeAttachments, nil
}

// ListVolAttachments ...
func (f *ServerClient) ListVolAttachments(serverid string) (models.GetServersServerIDVolumeAttachmentsOKBodyVolumeAttachments, error) {
	return f.ListVolAttachmentsWithFilter(serverid, "", "")
}

// GetVolAttachment ...
func (f *ServerClient) GetVolAttachment(serverid, volAttachID string) (*models.ServerVolumeAttachment, error) {
	params := compute.NewGetServersServerIDVolumeAttachmentsIDParams().
		WithServerID(serverid).
		WithID(volAttachID)
	resp, err := f.session.Riaas.Compute.GetServersServerIDVolumeAttachmentsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// AttachVolume ...
func (f *ServerClient) AttachVolume(serverid, volumeID, name string, resourcegroupID string, tags []string) (*models.ServerVolumeAttachment, error) {

	body := models.PostServersServerIDVolumeAttachmentsParamsBody{}
	if name != "" {
		body.Name = name
	}
	body.Volume = &models.PostServersServerIDVolumeAttachmentsParamsBodyVolume{
		ID: strfmt.UUID(volumeID),
	}

	if resourcegroupID != "" {
		rgref := models.PostServersServerIDVolumeAttachmentsParamsBodyResourceGroup{
			ID: strfmt.UUID(resourcegroupID),
		}
		body.ResourceGroup = &rgref
	}

	if len(tags) != 0 {
		body.Tags = tags
	}

	params := compute.NewPostServersServerIDVolumeAttachmentsParams().WithServerID(serverid).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PostServersServerIDVolumeAttachments(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteVolAttachment ...
func (f *ServerClient) DeleteVolAttachment(serverid, volAttachID string) error {
	params := compute.NewDeleteServersServerIDVolumeAttachmentsIDParams().WithServerID(serverid).WithID(volAttachID)
	_, err := f.session.Riaas.Compute.DeleteServersServerIDVolumeAttachmentsID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// UpdateVolAttachment ...
func (f *ServerClient) UpdateVolAttachment(serverid, volAttachID, name string) (*models.ServerVolumeAttachment, error) {

	body := models.PatchServersServerIDVolumeAttachmentsIDParamsBody{}
	if name != "" {
		body.Name = name
	}

	params := compute.NewPatchServersServerIDVolumeAttachmentsIDParams().WithServerID(serverid).WithID(volAttachID).WithBody(&body)
	resp, err := f.session.Riaas.Compute.PatchServersServerIDVolumeAttachmentsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

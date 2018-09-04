package compute

import (
	"github.ibm.com/riaas/rias-api/riaas/client/compute"
	"github.ibm.com/riaas/rias-api/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// ImageClient ...
type ImageClient struct {
	session *session.Session
}

// NewImageClient ...
func NewImageClient(sess *session.Session) *ImageClient {
	return &ImageClient{
		sess,
	}
}

// ListWithFilter ...
func (f *ImageClient) ListWithFilter(tag, visibility, start string) (models.GetImagesOKBodyImages, string, error) {
	params := compute.NewGetImagesParams()
	if tag != "" {
		params = params.WithTag(&tag)
	}
	if visibility != "" {
		params = params.WithVisibility(&visibility)
	}

	if start != "" {
		params = params.WithStart(&start)
	}

	resp, err := f.session.Riaas.Compute.GetImages(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	return resp.Payload.Images, utils.GetNext(resp.Payload.Next), nil
}

// List ...
func (f *ImageClient) List(start string) (models.GetImagesOKBodyImages, string, error) {
	return f.ListWithFilter("", "", start)
}

// Get ...
func (f *ImageClient) Get(id string) (*models.Image, error) {
	params := compute.NewGetImagesIDParams().WithID(id)
	resp, err := f.session.Riaas.Compute.GetImagesID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

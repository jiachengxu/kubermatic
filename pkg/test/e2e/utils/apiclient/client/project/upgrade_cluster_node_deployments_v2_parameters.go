// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// NewUpgradeClusterNodeDeploymentsV2Params creates a new UpgradeClusterNodeDeploymentsV2Params object
// with the default values initialized.
func NewUpgradeClusterNodeDeploymentsV2Params() *UpgradeClusterNodeDeploymentsV2Params {
	var ()
	return &UpgradeClusterNodeDeploymentsV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeClusterNodeDeploymentsV2ParamsWithTimeout creates a new UpgradeClusterNodeDeploymentsV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeClusterNodeDeploymentsV2ParamsWithTimeout(timeout time.Duration) *UpgradeClusterNodeDeploymentsV2Params {
	var ()
	return &UpgradeClusterNodeDeploymentsV2Params{

		timeout: timeout,
	}
}

// NewUpgradeClusterNodeDeploymentsV2ParamsWithContext creates a new UpgradeClusterNodeDeploymentsV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeClusterNodeDeploymentsV2ParamsWithContext(ctx context.Context) *UpgradeClusterNodeDeploymentsV2Params {
	var ()
	return &UpgradeClusterNodeDeploymentsV2Params{

		Context: ctx,
	}
}

// NewUpgradeClusterNodeDeploymentsV2ParamsWithHTTPClient creates a new UpgradeClusterNodeDeploymentsV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeClusterNodeDeploymentsV2ParamsWithHTTPClient(client *http.Client) *UpgradeClusterNodeDeploymentsV2Params {
	var ()
	return &UpgradeClusterNodeDeploymentsV2Params{
		HTTPClient: client,
	}
}

/*UpgradeClusterNodeDeploymentsV2Params contains all the parameters to send to the API endpoint
for the upgrade cluster node deployments v2 operation typically these are written to a http.Request
*/
type UpgradeClusterNodeDeploymentsV2Params struct {

	/*Body*/
	Body *models.MasterVersion
	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) WithTimeout(timeout time.Duration) *UpgradeClusterNodeDeploymentsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) WithContext(ctx context.Context) *UpgradeClusterNodeDeploymentsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) WithHTTPClient(client *http.Client) *UpgradeClusterNodeDeploymentsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) WithBody(body *models.MasterVersion) *UpgradeClusterNodeDeploymentsV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) SetBody(body *models.MasterVersion) {
	o.Body = body
}

// WithClusterID adds the clusterID to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) WithClusterID(clusterID string) *UpgradeClusterNodeDeploymentsV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) WithProjectID(projectID string) *UpgradeClusterNodeDeploymentsV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the upgrade cluster node deployments v2 params
func (o *UpgradeClusterNodeDeploymentsV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeClusterNodeDeploymentsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
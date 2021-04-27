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
)

// NewGetAlertmanagerParams creates a new GetAlertmanagerParams object
// with the default values initialized.
func NewGetAlertmanagerParams() *GetAlertmanagerParams {
	var ()
	return &GetAlertmanagerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertmanagerParamsWithTimeout creates a new GetAlertmanagerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAlertmanagerParamsWithTimeout(timeout time.Duration) *GetAlertmanagerParams {
	var ()
	return &GetAlertmanagerParams{

		timeout: timeout,
	}
}

// NewGetAlertmanagerParamsWithContext creates a new GetAlertmanagerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAlertmanagerParamsWithContext(ctx context.Context) *GetAlertmanagerParams {
	var ()
	return &GetAlertmanagerParams{

		Context: ctx,
	}
}

// NewGetAlertmanagerParamsWithHTTPClient creates a new GetAlertmanagerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAlertmanagerParamsWithHTTPClient(client *http.Client) *GetAlertmanagerParams {
	var ()
	return &GetAlertmanagerParams{
		HTTPClient: client,
	}
}

/*GetAlertmanagerParams contains all the parameters to send to the API endpoint
for the get alertmanager operation typically these are written to a http.Request
*/
type GetAlertmanagerParams struct {

	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get alertmanager params
func (o *GetAlertmanagerParams) WithTimeout(timeout time.Duration) *GetAlertmanagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alertmanager params
func (o *GetAlertmanagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alertmanager params
func (o *GetAlertmanagerParams) WithContext(ctx context.Context) *GetAlertmanagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alertmanager params
func (o *GetAlertmanagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alertmanager params
func (o *GetAlertmanagerParams) WithHTTPClient(client *http.Client) *GetAlertmanagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alertmanager params
func (o *GetAlertmanagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get alertmanager params
func (o *GetAlertmanagerParams) WithClusterID(clusterID string) *GetAlertmanagerParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get alertmanager params
func (o *GetAlertmanagerParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the get alertmanager params
func (o *GetAlertmanagerParams) WithProjectID(projectID string) *GetAlertmanagerParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get alertmanager params
func (o *GetAlertmanagerParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertmanagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

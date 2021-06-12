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

// NewListClusterTemplatesParams creates a new ListClusterTemplatesParams object
// with the default values initialized.
func NewListClusterTemplatesParams() *ListClusterTemplatesParams {
	var ()
	return &ListClusterTemplatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListClusterTemplatesParamsWithTimeout creates a new ListClusterTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListClusterTemplatesParamsWithTimeout(timeout time.Duration) *ListClusterTemplatesParams {
	var ()
	return &ListClusterTemplatesParams{

		timeout: timeout,
	}
}

// NewListClusterTemplatesParamsWithContext creates a new ListClusterTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListClusterTemplatesParamsWithContext(ctx context.Context) *ListClusterTemplatesParams {
	var ()
	return &ListClusterTemplatesParams{

		Context: ctx,
	}
}

// NewListClusterTemplatesParamsWithHTTPClient creates a new ListClusterTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListClusterTemplatesParamsWithHTTPClient(client *http.Client) *ListClusterTemplatesParams {
	var ()
	return &ListClusterTemplatesParams{
		HTTPClient: client,
	}
}

/*ListClusterTemplatesParams contains all the parameters to send to the API endpoint
for the list cluster templates operation typically these are written to a http.Request
*/
type ListClusterTemplatesParams struct {

	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list cluster templates params
func (o *ListClusterTemplatesParams) WithTimeout(timeout time.Duration) *ListClusterTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list cluster templates params
func (o *ListClusterTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list cluster templates params
func (o *ListClusterTemplatesParams) WithContext(ctx context.Context) *ListClusterTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list cluster templates params
func (o *ListClusterTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list cluster templates params
func (o *ListClusterTemplatesParams) WithHTTPClient(client *http.Client) *ListClusterTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list cluster templates params
func (o *ListClusterTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the list cluster templates params
func (o *ListClusterTemplatesParams) WithProjectID(projectID string) *ListClusterTemplatesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list cluster templates params
func (o *ListClusterTemplatesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListClusterTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	pipeline_model "github.com/kubeflow/pipelines/backend/api/go_http_client/pipeline_model"
)

// NewCreatePipelineVersionParams creates a new CreatePipelineVersionParams object
// with the default values initialized.
func NewCreatePipelineVersionParams() *CreatePipelineVersionParams {
	var ()
	return &CreatePipelineVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePipelineVersionParamsWithTimeout creates a new CreatePipelineVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePipelineVersionParamsWithTimeout(timeout time.Duration) *CreatePipelineVersionParams {
	var ()
	return &CreatePipelineVersionParams{

		timeout: timeout,
	}
}

// NewCreatePipelineVersionParamsWithContext creates a new CreatePipelineVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePipelineVersionParamsWithContext(ctx context.Context) *CreatePipelineVersionParams {
	var ()
	return &CreatePipelineVersionParams{

		Context: ctx,
	}
}

// NewCreatePipelineVersionParamsWithHTTPClient creates a new CreatePipelineVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePipelineVersionParamsWithHTTPClient(client *http.Client) *CreatePipelineVersionParams {
	var ()
	return &CreatePipelineVersionParams{
		HTTPClient: client,
	}
}

/*CreatePipelineVersionParams contains all the parameters to send to the API endpoint
for the create pipeline version operation typically these are written to a http.Request
*/
type CreatePipelineVersionParams struct {

	/*Body
	  ResourceReference inside PipelineVersion specifies the pipeline that this
	version belongs to.

	*/
	Body *pipeline_model.APIPipelineVersion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create pipeline version params
func (o *CreatePipelineVersionParams) WithTimeout(timeout time.Duration) *CreatePipelineVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pipeline version params
func (o *CreatePipelineVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pipeline version params
func (o *CreatePipelineVersionParams) WithContext(ctx context.Context) *CreatePipelineVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pipeline version params
func (o *CreatePipelineVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pipeline version params
func (o *CreatePipelineVersionParams) WithHTTPClient(client *http.Client) *CreatePipelineVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pipeline version params
func (o *CreatePipelineVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create pipeline version params
func (o *CreatePipelineVersionParams) WithBody(body *pipeline_model.APIPipelineVersion) *CreatePipelineVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create pipeline version params
func (o *CreatePipelineVersionParams) SetBody(body *pipeline_model.APIPipelineVersion) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePipelineVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
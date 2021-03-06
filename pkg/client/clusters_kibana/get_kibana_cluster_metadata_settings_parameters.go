// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package clusters_kibana

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

// NewGetKibanaClusterMetadataSettingsParams creates a new GetKibanaClusterMetadataSettingsParams object
// with the default values initialized.
func NewGetKibanaClusterMetadataSettingsParams() *GetKibanaClusterMetadataSettingsParams {
	var ()
	return &GetKibanaClusterMetadataSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKibanaClusterMetadataSettingsParamsWithTimeout creates a new GetKibanaClusterMetadataSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKibanaClusterMetadataSettingsParamsWithTimeout(timeout time.Duration) *GetKibanaClusterMetadataSettingsParams {
	var ()
	return &GetKibanaClusterMetadataSettingsParams{

		timeout: timeout,
	}
}

// NewGetKibanaClusterMetadataSettingsParamsWithContext creates a new GetKibanaClusterMetadataSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKibanaClusterMetadataSettingsParamsWithContext(ctx context.Context) *GetKibanaClusterMetadataSettingsParams {
	var ()
	return &GetKibanaClusterMetadataSettingsParams{

		Context: ctx,
	}
}

// NewGetKibanaClusterMetadataSettingsParamsWithHTTPClient creates a new GetKibanaClusterMetadataSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKibanaClusterMetadataSettingsParamsWithHTTPClient(client *http.Client) *GetKibanaClusterMetadataSettingsParams {
	var ()
	return &GetKibanaClusterMetadataSettingsParams{
		HTTPClient: client,
	}
}

/*GetKibanaClusterMetadataSettingsParams contains all the parameters to send to the API endpoint
for the get kibana cluster metadata settings operation typically these are written to a http.Request
*/
type GetKibanaClusterMetadataSettingsParams struct {

	/*ClusterID
	  The Kibana deployment identifier.

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kibana cluster metadata settings params
func (o *GetKibanaClusterMetadataSettingsParams) WithTimeout(timeout time.Duration) *GetKibanaClusterMetadataSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kibana cluster metadata settings params
func (o *GetKibanaClusterMetadataSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kibana cluster metadata settings params
func (o *GetKibanaClusterMetadataSettingsParams) WithContext(ctx context.Context) *GetKibanaClusterMetadataSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kibana cluster metadata settings params
func (o *GetKibanaClusterMetadataSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kibana cluster metadata settings params
func (o *GetKibanaClusterMetadataSettingsParams) WithHTTPClient(client *http.Client) *GetKibanaClusterMetadataSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kibana cluster metadata settings params
func (o *GetKibanaClusterMetadataSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get kibana cluster metadata settings params
func (o *GetKibanaClusterMetadataSettingsParams) WithClusterID(clusterID string) *GetKibanaClusterMetadataSettingsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get kibana cluster metadata settings params
func (o *GetKibanaClusterMetadataSettingsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKibanaClusterMetadataSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

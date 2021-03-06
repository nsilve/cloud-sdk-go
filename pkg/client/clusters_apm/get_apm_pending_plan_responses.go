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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetApmPendingPlanReader is a Reader for the GetApmPendingPlan structure.
type GetApmPendingPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApmPendingPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApmPendingPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApmPendingPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewGetApmPendingPlanPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApmPendingPlanOK creates a GetApmPendingPlanOK with default headers values
func NewGetApmPendingPlanOK() *GetApmPendingPlanOK {
	return &GetApmPendingPlanOK{}
}

/*GetApmPendingPlanOK handles this case with default header values.

The pending plan is applied to the APM server.
*/
type GetApmPendingPlanOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.ApmPlan
}

func (o *GetApmPendingPlanOK) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/plan/pending][%d] getApmPendingPlanOK  %+v", 200, o.Payload)
}

func (o *GetApmPendingPlanOK) GetPayload() *models.ApmPlan {
	return o.Payload
}

func (o *GetApmPendingPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.ApmPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApmPendingPlanNotFound creates a GetApmPendingPlanNotFound with default headers values
func NewGetApmPendingPlanNotFound() *GetApmPendingPlanNotFound {
	return &GetApmPendingPlanNotFound{}
}

/*GetApmPendingPlanNotFound handles this case with default header values.

The {cluster_id} can't be found. (code: 'clusters.cluster_not_found')
*/
type GetApmPendingPlanNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *GetApmPendingPlanNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/plan/pending][%d] getApmPendingPlanNotFound  %+v", 404, o.Payload)
}

func (o *GetApmPendingPlanNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetApmPendingPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApmPendingPlanPreconditionFailed creates a GetApmPendingPlanPreconditionFailed with default headers values
func NewGetApmPendingPlanPreconditionFailed() *GetApmPendingPlanPreconditionFailed {
	return &GetApmPendingPlanPreconditionFailed{}
}

/*GetApmPendingPlanPreconditionFailed handles this case with default header values.

The APM server is unable to finish provisioning, or the provisioning failed. Apply a plan, then try again. (code: 'clusters.cluster_plan_state_error')
*/
type GetApmPendingPlanPreconditionFailed struct {
	Payload *models.BasicFailedReply
}

func (o *GetApmPendingPlanPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/plan/pending][%d] getApmPendingPlanPreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetApmPendingPlanPreconditionFailed) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetApmPendingPlanPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

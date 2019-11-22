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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// StartDeploymentResourceInstancesAllReader is a Reader for the StartDeploymentResourceInstancesAll structure.
type StartDeploymentResourceInstancesAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDeploymentResourceInstancesAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewStartDeploymentResourceInstancesAllAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewStartDeploymentResourceInstancesAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStartDeploymentResourceInstancesAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewStartDeploymentResourceInstancesAllUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewStartDeploymentResourceInstancesAllRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStartDeploymentResourceInstancesAllAccepted creates a StartDeploymentResourceInstancesAllAccepted with default headers values
func NewStartDeploymentResourceInstancesAllAccepted() *StartDeploymentResourceInstancesAllAccepted {
	return &StartDeploymentResourceInstancesAllAccepted{}
}

/*StartDeploymentResourceInstancesAllAccepted handles this case with default header values.

The start command was issued successfully
*/
type StartDeploymentResourceInstancesAllAccepted struct {
	Payload models.DeploymentResourceCommandResponse
}

func (o *StartDeploymentResourceInstancesAllAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllAccepted  %+v", 202, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDeploymentResourceInstancesAllForbidden creates a StartDeploymentResourceInstancesAllForbidden with default headers values
func NewStartDeploymentResourceInstancesAllForbidden() *StartDeploymentResourceInstancesAllForbidden {
	return &StartDeploymentResourceInstancesAllForbidden{}
}

/*StartDeploymentResourceInstancesAllForbidden handles this case with default header values.

The start command was prohibited for the given Resource.
*/
type StartDeploymentResourceInstancesAllForbidden struct {
	Payload *models.BasicFailedReply
}

func (o *StartDeploymentResourceInstancesAllForbidden) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllForbidden  %+v", 403, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDeploymentResourceInstancesAllNotFound creates a StartDeploymentResourceInstancesAllNotFound with default headers values
func NewStartDeploymentResourceInstancesAllNotFound() *StartDeploymentResourceInstancesAllNotFound {
	return &StartDeploymentResourceInstancesAllNotFound{}
}

/*StartDeploymentResourceInstancesAllNotFound handles this case with default header values.

The Resource specified by {ref_id} cannot be found
*/
type StartDeploymentResourceInstancesAllNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *StartDeploymentResourceInstancesAllNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllNotFound  %+v", 404, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDeploymentResourceInstancesAllUnprocessableEntity creates a StartDeploymentResourceInstancesAllUnprocessableEntity with default headers values
func NewStartDeploymentResourceInstancesAllUnprocessableEntity() *StartDeploymentResourceInstancesAllUnprocessableEntity {
	return &StartDeploymentResourceInstancesAllUnprocessableEntity{}
}

/*StartDeploymentResourceInstancesAllUnprocessableEntity handles this case with default header values.

The command sent to a Resource found the Resource in an illegal state, the error message gives more details
*/
type StartDeploymentResourceInstancesAllUnprocessableEntity struct {
	Payload *models.BasicFailedReply
}

func (o *StartDeploymentResourceInstancesAllUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDeploymentResourceInstancesAllRetryWith creates a StartDeploymentResourceInstancesAllRetryWith with default headers values
func NewStartDeploymentResourceInstancesAllRetryWith() *StartDeploymentResourceInstancesAllRetryWith {
	return &StartDeploymentResourceInstancesAllRetryWith{}
}

/*StartDeploymentResourceInstancesAllRetryWith handles this case with default header values.

elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type StartDeploymentResourceInstancesAllRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *StartDeploymentResourceInstancesAllRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllRetryWith  %+v", 449, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
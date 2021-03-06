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

package deploymentapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestUpdate(t *testing.T) {
	type args struct {
		params UpdateParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentUpdateResponse
		err  string
	}{
		{
			name: "fails on parameter validation",
			err: multierror.NewPrefixed("deployment update",
				apierror.ErrMissingAPI,
				errors.New("request payload cannot be empty"),
				apierror.ErrDeploymentID,
			).Error(),
		},
		{
			name: "fails on API error",
			args: args{params: UpdateParams{
				DeploymentID: mock.ValidClusterID,
				API:          api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				Request:      &models.DeploymentUpdateRequest{},
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "succeeds updating to 7.4.1",
			args: args{params: UpdateParams{
				DeploymentID: mock.ValidClusterID,
				API: api.NewMock(mock.New200Response(mock.NewStructBody(models.DeploymentUpdateResponse{
					ID:   ec.String(mock.ValidClusterID),
					Name: ec.String("my example cluster"),
					Resources: []*models.DeploymentResource{
						{
							ID:     ec.String("01fa285876f54e699da3d3d6fd8a84f1"),
							Kind:   ec.String("elasticsearch"),
							RefID:  ec.String("my-es-cluster"),
							Region: ec.String("ece-region"),
						},
						{
							ElasticsearchClusterRefID: "my-es-cluster",
							ID:                        ec.String("ca8ac6555f0d43d8ba1048c98ea60265"),
							Kind:                      ec.String("kibana"),
							RefID:                     ec.String("my-kibana-instance"),
							Region:                    ec.String("ece-region"),
						},
					},
				}))),
				Request: &models.DeploymentUpdateRequest{
					Name: "my example cluster",
					Resources: &models.DeploymentUpdateResources{
						Elasticsearch: []*models.ElasticsearchPayload{
							{Plan: &models.ElasticsearchClusterPlan{
								Elasticsearch: &models.ElasticsearchConfiguration{Version: "7.4.1"},
							}},
						},
						Kibana: []*models.KibanaPayload{
							{Plan: &models.KibanaClusterPlan{
								Kibana: &models.KibanaConfiguration{Version: "7.4.1"},
							}},
						},
					},
				},
			}},
			want: &models.DeploymentUpdateResponse{
				ID:   ec.String(mock.ValidClusterID),
				Name: ec.String("my example cluster"),
				Resources: []*models.DeploymentResource{
					{
						ID:     ec.String("01fa285876f54e699da3d3d6fd8a84f1"),
						Kind:   ec.String("elasticsearch"),
						RefID:  ec.String("my-es-cluster"),
						Region: ec.String("ece-region"),
					},
					{
						ElasticsearchClusterRefID: "my-es-cluster",
						ID:                        ec.String("ca8ac6555f0d43d8ba1048c98ea60265"),
						Kind:                      ec.String("kibana"),
						RefID:                     ec.String("my-kibana-instance"),
						Region:                    ec.String("ece-region"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Update(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Update() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

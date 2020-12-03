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

package eskeystoreapi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestGet(t *testing.T) {
	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.KeystoreContents
		err  string
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid elasticsearch keystore get params",
				apierror.ErrMissingAPI,
				apierror.ErrDeploymentID,
			).Error(),
		},
		{
			name: "succeeds",
			args: args{params: GetParams{
				DeploymentID: mock.ValidClusterID,
				RefID:        "main-elasticsearch",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/elasticsearch/main-elasticsearch/keystore",
					},
					mock.NewStructBody(models.KeystoreContents{
						Secrets: map[string]models.KeystoreSecret{
							"some-secret":       {Value: "some-value"},
							"some-other-secret": {AsFile: ec.Bool(true), Value: "some-file"},
						},
					}),
				)),
			}},
			want: &models.KeystoreContents{
				Secrets: map[string]models.KeystoreSecret{
					"some-secret":       {Value: "some-value"},
					"some-other-secret": {AsFile: ec.Bool(true), Value: "some-file"},
				},
			},
		},
		{
			name: "succeeds with RefID discovery",
			args: args{params: GetParams{
				DeploymentID: mock.ValidClusterID,
				API: api.NewMock(
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed",
							Query: url.Values{
								"convert_legacy_plans": {"false"},
								"enrich_with_template": {"true"},
								"show_metadata":        {"false"},
								"show_plan_defaults":   {"false"},
								"show_plan_history":    {"false"},
								"show_plan_logs":       {"false"},
								"show_plans":           {"false"},
								"show_security":        {"false"},
								"show_settings":        {"false"},
								"show_system_alerts":   {"5"},
							},
						},
						mock.NewStructBody(models.DeploymentGetResponse{
							Healthy: ec.Bool(true),
							ID:      ec.String(mock.ValidClusterID),
							Resources: &models.DeploymentResources{
								Elasticsearch: []*models.ElasticsearchResourceInfo{{
									ID:    ec.String(mock.ValidClusterID),
									RefID: ec.String("elasticsearch"),
								}},
							},
						}),
					),
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/elasticsearch/elasticsearch/keystore",
						},
						mock.NewStructBody(models.KeystoreContents{
							Secrets: map[string]models.KeystoreSecret{
								"some-secret":       {Value: "some-value"},
								"some-other-secret": {AsFile: ec.Bool(true), Value: "some-file"},
							},
						}),
					),
				),
			}},
			want: &models.KeystoreContents{
				Secrets: map[string]models.KeystoreSecret{
					"some-secret":       {Value: "some-value"},
					"some-other-secret": {AsFile: ec.Bool(true), Value: "some-file"},
				},
			},
		},
		{
			name: "fails on RefID discovery",
			args: args{params: GetParams{
				DeploymentID: mock.ValidClusterID,
				API: api.NewMock(
					mock.New500ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed",
							Query: url.Values{
								"convert_legacy_plans": {"false"},
								"enrich_with_template": {"true"},
								"show_metadata":        {"false"},
								"show_plan_defaults":   {"false"},
								"show_plan_history":    {"false"},
								"show_plan_logs":       {"false"},
								"show_plans":           {"false"},
								"show_security":        {"false"},
								"show_settings":        {"false"},
								"show_system_alerts":   {"5"},
							},
						},
						mock.SampleInternalError().Response.Body,
					),
				),
			}},
			err: mock.MultierrorInternalError.Error(),
		},
		{
			name: "fails on API error",
			args: args{params: GetParams{
				DeploymentID: mock.ValidClusterID,
				RefID:        "main-elasticsearch",
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/elasticsearch/main-elasticsearch/keystore",
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

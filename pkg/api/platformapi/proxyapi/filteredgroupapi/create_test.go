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

package filteredgroupapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestCreate(t *testing.T) {
	var proxiesFilteredGroup = `
	{
      "expected_proxies_count": 15,
      "filters": [
        {
          "key": "proxyType",
          "value": "main-nextgen"
        }
      ],
      "id": "test2"
	}`
	type args struct {
		params CreateParams
	}
	tests := []struct {
		name string
		args args
		want *models.ProxiesFilteredGroup
		err  string
	}{
		{
			name: "Proxies filtered group create succeeds",
			args: args{params: CreateParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body:       ioutil.NopCloser(strings.NewReader(proxiesFilteredGroup)),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/proxies/filtered-groups",
						Body:   mock.NewStringBody(`{"expected_proxies_count":15,"filters":[{"key":"proxyType","value":"main-nextgen"}],"id":"test2"}` + "\n"),
					},
				}),
				ID: "test2",
				Filters: map[string]string{
					"proxyType": "main-nextgen",
				},
				ExpectedProxiesCount: 15,
			}},
			want: &models.ProxiesFilteredGroup{
				ExpectedProxiesCount: ec.Int32(15),
				Filters: []*models.ProxiesFilter{
					{
						Key:   ec.String("proxyType"),
						Value: ec.String("main-nextgen"),
					},
				},
				ID: *ec.String("test2"),
			},
		},
		{
			name: "Proxies filtered group create fails with 403 Forbidden",
			args: args{params: CreateParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{Response: http.Response{
					StatusCode: http.StatusForbidden,
					Status:     http.StatusText(http.StatusForbidden),
					Body:       mock.NewStringBody(`{"error": "some error"}`),
				}}),
				ID: "test2",
				Filters: map[string]string{
					"proxyType": "main-nextgen",
				},
				ExpectedProxiesCount: 15,
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "Proxies filtered group create fails due validation",
			args: args{params: CreateParams{}},
			err: multierror.NewPrefixed("invalid filtered group params",
				errors.New("id is not specified and is required for the operation"),
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
				errors.New("filters is not specified and is required for the operation"),
				errors.New("expected proxies count must be greater than 0"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}

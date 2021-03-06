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

package snaprepoapi

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParseGenericConfig(t *testing.T) {
	var invalidReaderContent = `such invalid {}content`
	var simpleJSONConfig = `{"a_key": "avalue"}`
	var simpleYAMLConfig = `
akey: avalue
anotherkey: anothervalue
`[1:]

	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    GenericConfig
		wantErr bool
	}{
		{
			name: "returns an error if the reader is nil",
			args: args{
				input: nil,
			},
			wantErr: true,
		},
		{
			name: "returns an error if the reader's content is not a valid yaml|json",
			args: args{
				input: strings.NewReader(invalidReaderContent),
			},
			wantErr: true,
		},
		{
			name: "Correctly reads a JSON formatted reader",
			args: args{
				input: strings.NewReader(simpleJSONConfig),
			},
			want: GenericConfig{
				"a_key": "avalue",
			},
			wantErr: false,
		},
		{
			name: "Correctly reads a YAML formatted reader",
			args: args{
				input: strings.NewReader(simpleYAMLConfig),
			},
			want: GenericConfig{
				"akey":       "avalue",
				"anotherkey": "anothervalue",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGenericConfig(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseGenericConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseGenericConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

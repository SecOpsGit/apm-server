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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("apm-server", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}

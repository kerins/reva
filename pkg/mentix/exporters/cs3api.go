// Copyright 2018-2020 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package exporters

import (
	"github.com/rs/zerolog"

	"github.com/cs3org/reva/pkg/mentix/config"
	"github.com/cs3org/reva/pkg/mentix/exporters/cs3api"
)

// CS3APIExporter implements the CS3API exporter.
type CS3APIExporter struct {
	BaseRequestExporter
}

// Activate activates the exporter.
func (exporter *CS3APIExporter) Activate(conf *config.Configuration, log *zerolog.Logger) error {
	if err := exporter.BaseExporter.Activate(conf, log); err != nil {
		return err
	}

	// Store CS3API specifics
	exporter.endpoint = conf.CS3API.Endpoint
	exporter.defaultMethodHandler = cs3api.HandleDefaultQuery

	return nil
}

// GetName returns the display name of the exporter.
func (exporter *CS3APIExporter) GetName() string {
	return "CS3API"
}

func init() {
	registerExporter(config.ExporterIDCS3API, &CS3APIExporter{})
}

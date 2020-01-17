// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filter_test

import (
	"testing"

	"github.com/banzaicloud/logging-operator/pkg/sdk/model/filter"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/render"
	"github.com/ghodss/yaml"
)

func TestDedot(t *testing.T) {
	CONFIG := []byte(`
de_dot_separator: "-"
de_dot_nested: true
n_lines: 10

`)
	expected := `
<filter **>
  @type dedot
  @id test_dedot
  de_dot_nested true
  de_dot_separator -
</filter>
`
	parser := &filter.DedotFilterConfig{}
	yaml.Unmarshal(CONFIG, parser)
	test := render.NewOutputPluginTest(t, parser)
	test.DiffResult(expected)
}
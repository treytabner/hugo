// Copyright 2017 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tpl

import (
	"html/template"

	"github.com/eknkc/amber"
)

func (gt *GoHTMLTemplate) CompileAmberWithTemplate(b []byte, path string, t *template.Template) (*template.Template, error) {
	c := amber.New()

	if err := c.ParseData(b, path); err != nil {
		return nil, err
	}

	data, err := c.CompileString()

	if err != nil {
		return nil, err
	}

	tpl, err := t.Funcs(gt.amberFuncMap).Parse(data)

	if err != nil {
		return nil, err
	}

	return tpl, nil
}

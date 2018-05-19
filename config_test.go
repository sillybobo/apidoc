// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"testing"

	"github.com/issue9/assert"

	"github.com/caixw/apidoc/input"
	"github.com/caixw/apidoc/openapi"
	"github.com/caixw/apidoc/output"
)

func TestConfig_sanitize(t *testing.T) {
	a := assert.New(t)

	conf := &config{}
	err := conf.sanitize()
	a.Error(err)
	a.Equal(err.(*openapi.Error).Field, "version")

	// 版本号错误
	conf.Version = "4.0"
	err = conf.sanitize()
	a.Error(err)
	a.Equal(err.(*openapi.Error).Field, "version")

	// 未声明 inputs
	conf.Version = "4.0.1"
	err = conf.sanitize()
	a.Error(err)
	a.Equal(err.(*openapi.Error).Field, "inputs")

	// 未声明 output
	conf.Inputs = []*input.Options{{}}
	err = conf.sanitize()
	a.Error(err)
	a.Equal(err.(*openapi.Error).Field, "output")

	// 查看错误提示格式是否正确
	conf.Output = &output.Options{}
	conf.Inputs = append(conf.Inputs, &input.Options{
		Lang: "123",
	})
	err = conf.sanitize()
	a.Error(err)
	a.True(strings.HasPrefix(err.(*openapi.Error).Field, "inputs[0]"))
}

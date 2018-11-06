// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package doc

import (
	"bytes"
	"testing"

	"github.com/issue9/assert"
)

func TestBody_parseExample(t *testing.T) {
	a := assert.New(t)
	body := &Body{}

	tag := newTag(`@apiExample application/json summary text
{
	"id": 1,
	"name": "name"
}`)
	body.parseExample(tag)
	e := body.Examples[0]
	a.Equal(e.Mimetype, "application/json").
		Equal(e.Summary, "summary text").
		Equal(e.Value, `{
	"id": 1,
	"name": "name"
}`)

	// 长度不够
	tag = newTag("application/json")
	body.parseExample(tag)
}

func TestBody_parseHeader(t *testing.T) {
	a := assert.New(t)
	body := &Body{}

	tag := newTag(`@apiExample content-type required json 或是 xml`)
	body.parseHeader(tag)
	h := body.Headers[0]
	a.Equal(h.Summary, "json 或是 xml").
		Equal(h.Name, "content-type").
		False(h.Optional)

	tag = newTag(`@apiExample ETag optional etag`)
	body.parseHeader(tag)
	h = body.Headers[1]
	a.Equal(h.Summary, "etag").
		Equal(h.Name, "ETag").
		True(h.Optional)

	// 长度不够
	tag = newTag("ETag")
	body.parseHeader(tag)
}

func TestIsOptional(t *testing.T) {
	a := assert.New(t)

	a.False(isOptional(requiredBytes))
	a.False(isOptional(bytes.ToUpper(requiredBytes)))
	a.True(isOptional([]byte("optional")))
	a.True(isOptional([]byte("Optional")))
}

func TestNewResponse(t *testing.T) {
	a := assert.New(t)
	l := newLexer(`@apiHeader content-type optional 指定内容类型
	@apiParam id int required 唯一 ID
	@apiParam name string required 名称
	@apiParam nickname string optional 昵称
	@apiExample json 默认返回示例
	{
		"id": 1,
		"name": "name",
		"nickname": "nickname"
	}
	@apiUnknown xxx`)
	tag := newTag(`@apiResponse 200 array.object * 通用的返回内容定义`)

	resp, ok := newResponse(l, tag)
	a.True(ok).NotNil(resp)
	a.Equal(resp.Status, 200).
		Equal(resp.Mimetype, "*")
	a.Equal(len(resp.Headers), 1).
		Equal(resp.Headers[0].Name, "content-type").
		Equal(resp.Headers[0].Summary, "指定内容类型").
		True(resp.Headers[0].Optional)
	a.NotNil(resp.Type).
		Equal(resp.Type.Type, Array)
}

func TestResponses_parseResponse(t *testing.T) {
	a := assert.New(t)
	d := &responses{}

	l := newLexer(`@apiHeader content-type optional 指定内容类型
	@apiParam id int required 唯一 ID
	@apiParam name string required 名称
	@apiParam nickname string optional 昵称
	@apiExample json 默认返回示例
	{
		"id": 1,
		"name": "name",
		"nickname": "nickname"
	}
	@apiUnknown xxx`)
	tag := newTag(`@apiResponse 200 array.object * 通用的返回内容定义`)

	d.parseResponse(l, tag)
	a.Equal(len(d.Responses), 1)
	resp := d.Responses[0]
	a.Equal(resp.Status, 200).
		Equal(resp.Mimetype, "*").
		Equal(resp.Type.Description, "通用的返回内容定义")
	a.Equal(len(resp.Headers), 1).
		Equal(resp.Headers[0].Name, "content-type").
		Equal(resp.Headers[0].Summary, "指定内容类型").
		True(resp.Headers[0].Optional)
	a.NotNil(resp.Type).
		Equal(resp.Type.Type, Array)

	// 可以添加多次。
	d.parseResponse(l, tag)
	a.Equal(len(d.Responses), 2)
	resp = d.Responses[1]
	a.Equal(resp.Status, 200).
		Equal(resp.Mimetype, "*")

	// 忽略可选参数
	tag = newTag(`@apiResponse 200 array.object * `)
	d.parseResponse(l, tag)
	a.Equal(len(d.Responses), 3)
	resp = d.Responses[2]
	a.Equal(resp.Status, 200).
		Equal(resp.Mimetype, "*").
		Empty(resp.Type.Description)
}

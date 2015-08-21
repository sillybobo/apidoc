// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"github.com/caixw/apidoc/core"
	"github.com/caixw/apidoc/scanner"
)

// 各编程语言相关的参数
type lang struct {
	exts []string // 扩展名列表
	scan core.ScanFunc
}

var langs = map[string]*lang{
	"go": &lang{
		exts: []string{".go"},
		scan: scanner.CStyle,
	},

	"cpp": &lang{
		exts: []string{".h", ".cpp", ".cxx", ".c"},
		scan: scanner.CStyle,
	},

	"c": &lang{
		exts: []string{".h", ".c"},
		scan: scanner.CStyle,
	},

	"php": &lang{
		exts: []string{".php"},
		scan: scanner.CStyle,
	},

	"js": &lang{
		exts: []string{".js"},
		scan: scanner.CStyle,
	},
}

// 各扩展名对应的语言。
// 数据由init函数从上面的langs数据中分析获得。
var extsIndex = map[string]string{}

func init() {
	for k, lang := range langs {
		for _, ext := range lang.exts {
			extsIndex[ext] = k
		}
	}
}
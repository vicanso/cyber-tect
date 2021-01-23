// Copyright 2020 tree xie
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

package service

import (
	"bytes"
	"errors"
	"time"

	"github.com/felixge/fgprof"
)

type ProfSrv struct{}

func (srv *ProfSrv) Get(d time.Duration) (result *bytes.Buffer, err error) {
	// 禁止拉取超过1分钟的prof
	if d > 1*time.Minute {
		err = errors.New("duration should be less than 1m")
		return
	}
	result = &bytes.Buffer{}
	done := fgprof.Start(result, fgprof.FormatPprof)
	time.Sleep(d)
	err = done()
	if err != nil {
		return
	}
	return
}

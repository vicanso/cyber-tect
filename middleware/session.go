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

package middleware

import (
	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/elton"
	jwt "github.com/vicanso/elton-jwt"
)

// NewSession new session middleware
func NewSession() elton.Handler {
	scf := config.GetSessionConfig()
	ttlToken := &jwt.TTLToken{
		TTL: scf.TTL,
		// 密钥用于加密数据，需保密
		Secret: []byte(scf.Secret),
	}

	// 用于初始化创建token使用（此时可能token还没有或者已过期)
	return jwt.NewJWT(jwt.Config{
		CookieName:  scf.Key,
		TTLToken:    ttlToken,
		Passthrough: true,
	})
}

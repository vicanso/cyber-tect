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

package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfigENV(t *testing.T) {
	assert := assert.New(t)
	originENV := env
	defer func() {
		env = originENV
	}()

	env = "test"
	assert.Equal(env, GetENV())
}

func TestBasicConfig(t *testing.T) {
	assert := assert.New(t)

	basicConfig := GetBasicConfig()
	assert.Equal("cybertect", basicConfig.Name)
	assert.Equal(uint(1000), basicConfig.RequestLimit)
	assert.Equal(":7001", basicConfig.Listen)
}

func TestSessionConfig(t *testing.T) {
	assert := assert.New(t)

	sessionConfig := GetSessionConfig()
	assert.Equal(240*time.Hour, sessionConfig.TTL)
	assert.Equal("cybertect", sessionConfig.Key)
	assert.Equal("/", sessionConfig.CookiePath)
	assert.Equal([]string{"cuttlefish", "secret"}, sessionConfig.Keys)
	assert.Equal("jt", sessionConfig.TrackKey)
}

func TestRedisConfig(t *testing.T) {
	assert := assert.New(t)

	redisConfig := GetRedisConfig()
	assert.Equal("127.0.0.1:6379", redisConfig.Addr)
	assert.Equal("", redisConfig.Password)
	assert.Equal(0, redisConfig.DB)
	assert.Equal(200*time.Millisecond, redisConfig.Slow)
	assert.Equal(uint32(1000), redisConfig.MaxProcessing)
}

func TestMailConfig(t *testing.T) {
	assert := assert.New(t)

	mailConfig := GetMailConfig()
	assert.Equal("smtp.office365.com", mailConfig.Host)
	assert.Equal(587, mailConfig.Port)
	assert.Equal("tree.xie@outlook.com", mailConfig.User)
	assert.Equal("EMAIL_PASS", mailConfig.Password)
}

func TestInfluxdbConfig(t *testing.T) {
	assert := assert.New(t)

	influxdbConfig := GetInfluxdbConfig()
	assert.Equal("http://127.0.0.1:8086", influxdbConfig.URI)
	assert.Equal("cybertect", influxdbConfig.Bucket)
	assert.Equal("bigTree", influxdbConfig.Org)
	assert.Equal("YXFOYqPAxtF2wxfVMn-M168Y8rbAFhOsMDKCVWaStmEtFTqFX9KtcYDQk5ouwCNJY0iyW1KVE-VhmmAAhnfG5w==", influxdbConfig.Token)
	assert.Equal(uint(100), influxdbConfig.BatchSize)
	assert.Equal(10*time.Second, influxdbConfig.FlushInterval)
	assert.False(influxdbConfig.Disabled)
}

func TestAlarmConfig(t *testing.T) {
	assert := assert.New(t)

	alarmConfig := GetAlarmConfig()
	assert.Equal([]string{"tree.xie@outlook.com"}, alarmConfig.Receivers)
}

func TestGetPostgresConfig(t *testing.T) {
	assert := assert.New(t)

	postgresConfig := GetPostgresConfig()
	assert.NotNil(postgresConfig.URI)
}

func TestGetLocationConfig(t *testing.T) {
	assert := assert.New(t)

	locationConfig := GetLocationConfig()
	assert.Equal("https://ip.npmtrend.com", locationConfig.BaseURL)
	assert.Equal("location", locationConfig.Name)
	assert.Equal(3*time.Second, locationConfig.Timeout)
}

func TestGetMinioConfig(t *testing.T) {
	assert := assert.New(t)

	minioConfig := GetMinioConfig()
	assert.Equal("127.0.0.1:9000", minioConfig.Endpoint)
	assert.Equal("origin", minioConfig.AccessKeyID)
	assert.Equal("test123456", minioConfig.SecretAccessKey)
	assert.False(minioConfig.SSL)
}

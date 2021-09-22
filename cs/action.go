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
package cs

const (
	// ActionLogin login
	ActionLogin = "login"
	// ActionRegister register
	ActionRegister = "register"
	// ActionLogout logout
	ActionLogout = "logout"

	// ActionUserInfoUpdate update user info
	ActionUserInfoUpdate = "updateUserInfo"
	// ActionUserMeUpdate update my info
	ActionUserMeUpdate = "updateUserMe"

	// ActionConfigurationAdd add configuration
	ActionConfigurationAdd = "addConfiguration"
	// ActionConfigurationUpdate update configuration
	ActionConfigurationUpdate = "updateConfiguration"

	// ActionAdminCleanCache clean cache
	ActionAdminCleanCache = "cleanCache"

	// ActionDetectorHTTPAdd add http detector
	ActionDetectorHTTPAdd = "addHTTPDetector"
	// ActionDetectorHTTPUpdate update http detector
	ActionDetectorHTTPUpdate = "updateHTTPDetector"
	// ActionDetectorDNSAdd add dns detector
	ActionDetectorDNSAdd = "addDNSDetector"
	// ActionDetectorDNSUpdate update dns detector
	ActionDetectorDNSUpdate = "updateDNSDetector"
	// ActionDetectorTCPAdd add tcp detector
	ActionDetectorTCPAdd = "addTCPDetector"
	// ActionDetectorTCPUpdate update tcp detector
	ActionDetectorTCPUpdate = "updateTCPDetector"
	// ActionDetectorPingAdd add ping detector
	ActionDetectorPingAdd = "addPingDetector"
	// ActionDetectorPingUpdate update ping detector
	ActionDetectorPingUpdate = "updatePingDetector"
)

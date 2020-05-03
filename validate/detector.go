// Copyright 2019 tree xie
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

package validate

import (
	"github.com/go-playground/validator/v10"
)

func init() {
	Add("xDetectorStatus", func(fl validator.FieldLevel) bool {
		return isZero(fl) || isInInt(fl, []int{
			1,
			2,
		})
	})
	AddAlias("xDetectorDescription", "min=0,max=1000")

	AddAlias("xDNSServer", "ip|hostname_port")

	AddAlias("xDNSHostname", "hostname")

	Add("xTCPNetwork", func(fl validator.FieldLevel) bool {
		return isZero(fl) || isInString(fl, []string{
			"ip4:icmp",
			"tcp",
		})
	})
}

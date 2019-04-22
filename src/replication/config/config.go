// Copyright Project Harbor Authors
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

package config

var (
	// Config is the configuration
	Config *Configuration
)

// Configuration holds the configuration information for replication
type Configuration struct {
	CoreURL         string
	RegistryURL     string
	TokenServiceURL string
	JobserviceURL   string
	SecretKey       string
	// TODO consider to use a specified secret for replication
	CoreSecret       string
	JobserviceSecret string
}
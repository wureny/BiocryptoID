// Copyright 2024 BiocryptoID
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flags

import "flag"

var (
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSGATEWAYURL      string
)

func init() {
	flag.StringVar(&AWSAccessKeyID, "aws-access-key-id", "", "AWS Access Key ID")
	flag.StringVar(&AWSSecretAccessKey, "aws-secret-access-key", "", "AWS Secret Access Key")
	flag.StringVar(&AWSGATEWAYURL, "aws-gateway-url", "", "AWS Gateway URL")
}

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

package web

// 生物特征注册的结构体
type BiometricRegister struct {
	FacialInfo   []byte `json:"facial_info"`
	Name         string `json:"name"`
	Birth        string `json:"birth"`
	Sex          string `json:"sex"`
	SecurityCode string `json:"security_code"`
}

// 生物特征认证的结构体
type BiometricAuth struct {
	Code       string `json:"code"`
	DID        string `json:"did"`
	FacialInfo []byte `json:"facial_info"`
	Msg        string `json:"msg"`
}

// VO register(前端传来的数据)
type Register struct{}

// VO auth(前端传来的数据)
type Auth struct{}

func ConvertRegisterToBiometricRegister(register Register) BiometricRegister {
	return BiometricRegister{}
}

func ConvertAuthToBiometricAuth(auth Auth) BiometricAuth {
	return BiometricAuth{}
}

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

import (
	"biocryptoID/internal/service"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log/slog"
)

func HandlerRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	slog.Info("Processing request data for request %s\n", request.RequestContext.RequestID)
	slog.Info("Path: %s\n", request.Path)

	headres := map[string]string{"Content-Type": "text/plain"}
	switch request.Path {
	case "/register":
		//将request中的body解析成Register结构体，调用convertRegisterToBiometricRegister函数
		//从request中解析出Register结构体
		body := request.Body
		req := Register{}
		err := json.Unmarshal([]byte(body), &req)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Headers:    headres,
				Body:       "Bad Request",
				StatusCode: 400,
			}, nil
		}
		//将req转换成BiometricRegister
		biometricRegister := ConvertRegisterToBiometricRegister(req)
		err = service.HandleRegister(ctx, biometricRegister)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Headers:    headres,
				Body:       "Internal Server Error",
				StatusCode: 500,
			}, nil
		}
		return events.APIGatewayProxyResponse{
			Headers: headres,
			//TODO implement body
			Body:       "Success",
			StatusCode: 200,
		}, nil
	case "/auth":
		body := request.Body
		req := Auth{}
		err := json.Unmarshal([]byte(body), &req)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Headers:    headres,
				Body:       "Bad Request",
				StatusCode: 400,
			}, nil
		}
		biometricAuth := ConvertAuthToBiometricAuth(req)
		err = service.HandleAuth(ctx, biometricAuth)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Headers:    headres,
				Body:       "Internal Server Error",
				StatusCode: 500,
			}, nil
		}
		return events.APIGatewayProxyResponse{
			Headers:    headres,
			Body:       "Success",
			StatusCode: 200,
		}, nil
	default:
		return events.APIGatewayProxyResponse{
			Headers:    headres,
			Body:       "Not Found, try again",
			StatusCode: 404,
		}, nil
	}
}

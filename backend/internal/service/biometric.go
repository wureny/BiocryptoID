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

package service

import (
	"biocryptoID/flags"
	"biocryptoID/internal/domain"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
)

type BiometricService struct{}

func HandleRegister(ctx context.Context, register domain.BiometricRegister) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			flags.AWSAccessKeyID,
			flags.AWSSecretAccessKey,
			"",
		)),
	)
	if err != nil {
		return err
	}
	cli := rekognition.NewFromConfig(cfg)
	params := &rekognition.DetectFacesInput{
		Image:      nil,
		Attributes: nil,
	}
	_, err = cli.DetectFaces(context.TODO(), params)
	return nil
}

func HandleAuth(ctx context.Context, Auth domain.BiometricAuth) error {
	return nil
}

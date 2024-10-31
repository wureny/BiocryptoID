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
	"biocryptoID/internal/awsconfig"
	"biocryptoID/internal/domain"
	"biocryptoID/internal/repository"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

type BiometricService struct{}

func HandleRegister(ctx context.Context, register domain.BiometricRegister) error {
	cfg, err := awsconfig.NewAWSConfig()
	if err != nil {
		return err
	}
	cli := rekognition.NewFromConfig(cfg)

	image := types.Image{
		Bytes:    register.FacialInfo,
		S3Object: nil,
	}
	params := &rekognition.DetectFacesInput{
		Image:      &image,
		Attributes: nil,
	}

	output, err := cli.DetectFaces(context.TODO(), params)
	if err != nil {
		return err
	}
	info := domain.BioInfo{
		Age:         20,
		FaceDetails: output.FaceDetails,
	}
	// Upload to S3
	err = repository.UploadBiometric(info)
	if err != nil {
		return err
	}
	return nil
}

func HandleAuth(ctx context.Context, Auth domain.BiometricAuth) error {
	return nil
}

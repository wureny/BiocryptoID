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

package repository

import (
	"biocryptoID/internal/awsconfig"
	"biocryptoID/internal/domain"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BiometricRepository struct {
}

func UploadBiometric(info domain.BioInfo) error {
	cfg, err := awsconfig.NewAWSConfig()
	if err != nil {
		return err
	}
	cli := s3.NewFromConfig(cfg)
	keyMsg := "did"
	bucket := "sss"
	_, err = cli.PutObject(context.TODO(), &s3.PutObjectInput{
		//TODO: flag,did
		Bucket: &bucket,
		Key:    &keyMsg,
		Body:   info,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetBiometric() error {
	return nil
}

func UpdateBiometric() error {
	return nil
}

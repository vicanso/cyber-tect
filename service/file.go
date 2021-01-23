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

package service

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/elton"
)

type (
	FileSrv struct{}
	// UploadParams 上传文件参数
	UploadParams struct {
		// Bucket 文件保存的bucket
		Bucket string
		// Name 文件名
		Name string
		// Reader 文件内容的reader
		Reader io.Reader
		// Size 文件大小
		Size int64
		Opts minio.PutObjectOptions
	}
)

var defaultMinioClient = mustNewMinioClient()

// mustNewMinioClient 初始化minio client
func mustNewMinioClient() *minio.Client {
	minioConfig := config.GetMinioConfig()
	c, err := minio.New(minioConfig.Endpoint, &minio.Options{
		Secure: minioConfig.SSL,
		Creds:  credentials.NewStaticV4(minioConfig.AccessKeyID, minioConfig.SecretAccessKey, ""),
	})
	if err != nil {
		panic(err)
	}
	return c
}

// Upload 上传文件
func (srv *FileSrv) Upload(ctx context.Context, params UploadParams) (info minio.UploadInfo, err error) {
	return defaultMinioClient.PutObject(ctx, params.Bucket, params.Name, params.Reader, params.Size, params.Opts)
}

// Get 获取文件
func (srv *FileSrv) Get(ctx context.Context, bucket, filename string) (*minio.Object, error) {
	return defaultMinioClient.GetObject(ctx, bucket, filename, minio.GetObjectOptions{})
}

// GetData 获取文件内容及对应的http头
func (srv *FileSrv) GetData(ctx context.Context, bucket, filename string) (data []byte, header http.Header, err error) {
	object, err := srv.Get(ctx, bucket, filename)
	if err != nil {
		return
	}
	statsInfo, err := object.Stat()
	if err != nil {
		return
	}
	header = make(http.Header)
	header.Set(elton.HeaderETag, statsInfo.ETag)
	header.Set(elton.HeaderContentType, statsInfo.ContentType)

	data, err = ioutil.ReadAll(object)
	if err != nil {
		return
	}
	return
}

// ResizeImage 调整图片尺寸
func (srv *FileSrv) ResizeImage(reader io.Reader, imageType string, width, height int) (buffer *bytes.Buffer, err error) {
	var img image.Image
	switch imageType {
	default:
		img, _, err = image.Decode(reader)
	case "png":
		img, err = png.Decode(reader)
	case "jpg":
		fallthrough
	case "jpeg":
		img, err = jpeg.Decode(reader)
	}
	if err != nil {
		return
	}
	if width == 0 && height == 0 {
		width = img.Bounds().Dx()
		height = img.Bounds().Dy()
	}
	img = imaging.Resize(img, width, height, imaging.Lanczos)
	buffer = new(bytes.Buffer)
	switch imageType {
	default:
		err = jpeg.Encode(buffer, img, &jpeg.Options{
			Quality: 80,
		})
	case "png":
		err = png.Encode(buffer, img)
	}
	if err != nil {
		return
	}
	return
}

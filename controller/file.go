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

package controller

import (
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/vicanso/elton"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/service"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/cybertect/validate"
)

type fileCtrl struct{}

// 响应相关定义
type (
	// fileUploadResp 文件上传响应
	fileUploadResp struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
	}
)

// 参数相关定义
type (
	// fileUploadParams 文件上传参数
	fileUploadParams struct {
		Bucket string `json:"bucket" validate:"required,xFileBucket"`
	}
)

func init() {
	ctrl := fileCtrl{}

	g := router.NewGroup("/files")

	// 上传文件
	g.POST(
		"/v1",
		loadUserSession,
		shouldBeLogin,
		ctrl.upload,
	)
}

// upload 上传文件
func (*fileCtrl) upload(c *elton.Context) error {
	params := fileUploadParams{}
	err := validate.Do(&params, c.Query())
	if err != nil {
		return err
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()
	us := getUserSession(c)
	contentType := header.Header.Get("Content-Type")
	fileType := strings.Split(contentType, "/")[1]
	name := util.GenXID() + "." + fileType
	info, err := fileSrv.Upload(c.Context(), service.UploadParams{
		Bucket: params.Bucket,
		Name:   name,
		Reader: file,
		Size:   header.Size,
		Opts: minio.PutObjectOptions{
			ContentType: contentType,
			UserTags: map[string]string{
				"account": us.MustGetInfo().Account,
			},
		},
	})
	if err != nil {
		return err
	}
	c.Body = &fileUploadResp{
		Name: info.Key,
		Size: info.Size,
	}
	return nil
}

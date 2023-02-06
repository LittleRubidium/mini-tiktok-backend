// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/golang-jwt/jwt/v4"
	"mime/multipart"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/publish"
	pkg_consts "mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"
)

// Paras 文件类型的参数接收单独定义
type Paras struct {
	data  *multipart.FileHeader `form:"data"`
	token string                `form:"token"`
	title string                `form:"title"`
}

// DouyinPublishAction .
// @router /douyin/publish/action/ [POST]
func DouyinPublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Paras
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	file, _ := req.data.Open()
	defer file.Close()
	fileContent, _ := ReadFileContent(file)

	claims, _ := GetClaimsFromTokenString(req.token)
	// TODO: register jwt middleware

	err = rpc.PublishVideo(ctx, &publish.PublishVideoRequest{
		UserId: int64(claims[pkg_consts.UserIdKey].(float64)),
		Data:   fileContent,
		Title:  req.title,
	})

	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	SendResponse(c, errno.Success, utils.H{})
}

func GetClaimsFromTokenString(token string) (map[string]interface{}, error) {
	t, _ := mw.JwtMiddleware.ParseTokenString(token)
	claims := jwt.MapClaims{}
	for key, value := range t.Claims.(jwt.MapClaims) {
		claims[key] = value
	}
	return claims, nil
}

func ReadFileContent(file multipart.File) ([]byte, error) {
	const fileChunkSize = 1 * (1 << 20) // read 1MB each time.
	fileContent := make([]byte, 0, 0)

	for {
		fileChunk := make([]byte, fileChunkSize)
		n, _ := file.Read(fileChunk)
		if n == 0 {
			break
		}
		fileChunk = fileChunk[:n]
		fileContent = append(fileContent, fileChunk...)
	}

	return fileContent, nil
}

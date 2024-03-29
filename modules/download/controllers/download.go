/*
 * @Author: liziwei01
 * @Date: 2023-05-09 23:00:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-05 02:10:09
 * @Description: file content
 */
package controllers

import (
	"path/filepath"

	"github.com/liziwei01/gin-lib/library/env"
	"github.com/liziwei01/gin-lib/library/response"
	"github.com/liziwei01/gin-monitor-appui/modules/download/constants"
	downloadModel "github.com/liziwei01/gin-monitor-appui/modules/download/model"

	"github.com/gin-gonic/gin"
)

func Stream(ctx *gin.Context) {
	inputs, hasError := getStreamPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}

	absPath := filepath.Join(env.Default.RootDir(), constants.DATA_LOCATION, inputs.Path)
	rlvPath := filepath.Base(absPath)
	ctx.FileAttachment(absPath, rlvPath)
}

func getStreamPars(ctx *gin.Context) (downloadModel.StreamPars, bool) {
	var inputs downloadModel.StreamPars

	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}

	return inputs, false
}

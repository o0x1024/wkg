package env

import (
	"os"
	"wkg/pkg/libs/fileOp"

	"go.uber.org/zap"
)

func EnvCheck() {
	var err error

	if !fileOp.PathExists("log") {
		err = os.MkdirAll("log", os.ModePerm)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return
		}
	}

	if !fileOp.PathExists("tmp") {
		err = os.MkdirAll("tmp", os.ModePerm)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return
		}

	}

	if !fileOp.PathExists("upload/img") {
		err = os.MkdirAll("upload/img", os.ModePerm)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return
		}
	}

}

package env

import (
	"os"
	"wkg/pkg/libs/fileOp"

	"go.uber.org/zap"
)

func EnvCheck() {
	exitst, err := fileOp.PathExists("./conf/subfinder.yaml")
	if err != nil {
		zap.S().Error("%s", err.Error())
		return
	}
	if !exitst {
		zap.S().Error("./conf/subfinder.yaml not found.")
		return
	}

	ex, err := fileOp.PathExists("log")
	if err != nil {
		zap.S().Errorf("%s", err.Error())
	}
	if !ex {
		err = os.MkdirAll("log", os.ModePerm)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return
		}
	}

	ex, err = fileOp.PathExists("tmp")
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}
	if !ex {
		err = os.MkdirAll("tmp", os.ModePerm)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return
		}

	}

	ex, err = fileOp.PathExists("upload/img")
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}
	if !ex {
		err = os.MkdirAll("upload/img", os.ModePerm)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return
		}
	}

}

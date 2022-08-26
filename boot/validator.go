package boot

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	validator2 "go-gin-api/pkg/validator"
)

func initValidator() {
	validate := binding.Validator.Engine().(*validator.Validate)
	validator2.InitTrans(validate)
}

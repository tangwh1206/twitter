package common

import (
	"github.com/gin-gonic/gin"
	"github.com/tangwh1206/twitter/faults"
)

func Response(res interface{}, err error) gin.H {
	if err == nil {
		return gin.H{
			"code": faults.CodeSuccess,
			"msg":  "success",
			"data": res,
		}
	}
	ft, ok := err.(faults.Fault)
	if ok {
		return gin.H{
			"code": ft.Code(),
			"msg":  ft.Message(),
			"data": "",
		}
	} else {
		return gin.H{
			"code": faults.CodeInternelError,
			"msg":  faults.MsgInternalError,
			"data": "",
		}
	}

}



package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "net/http"
)

// path paramter with name details will mapped to Details
type URI struct {
   Details string `json:"name" uri:"details"`
}

func main() {
   engine:=gin.New()
// adding path params to router
   engine.GET("/test/:details", func(context *gin.Context) {
      uri:=URI{}
      // binding to URI
      if err:=context.BindUri(&uri);err!=nil{
         context.AbortWithError(http.StatusBadRequest,err)
         return
      }
      fmt.Println(uri)
      context.JSON(http.StatusAccepted,&uri)
   })
   engine.Run(":3000")
}

/*
Note: Please check https://github.com/msniranjan18/go-gin-basic/blob/master/ex2/main.go for more info.
*/

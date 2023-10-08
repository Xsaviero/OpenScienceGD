package handlers

import (
	"fdfd/types"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// w http.ResponseWriter, r *http.Request
func RegisterUser(c *gin.Context) {
	var userInfo types.User
	err := c.BindJSON(&userInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Print(err)
	}
}

// func AuthentificateUser(c *gin.Context) {
// 	sm , err :=
// }
//	func New(s *internal.Storage) gin.HandlerFunc {
//		return func(c *gin.Context) {
//			var userInfo types.User
//			err := c.BindJSON(&userInfo)
//			if err != nil {
//				log.Fatal(err)
//			}
//			s.SaveUser(userInfo)
//		}
//	}

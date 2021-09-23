package controllers

import (
	"errors"
	"strconv"
	// "github.com/gin-gonic/gin"
)

// does not include test for zero
func getIDParam(idParam string) (uint, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errors.New("id should be a number")
	}
	return uint(id), nil
}

// // Response object as HTTP response
// type Response struct {
// 	Code int         `json:"code"`
// 	Msg  string      `json:"msg"`
// 	Data interface{} `json:"data"`
// }
//
// // HTTPRes normalize HTTP Response format
// func HTTPRes(c *gin.Context, httpCode int, msg string, data interface{}) {
// 	c.JSON(httpCode, Response{
// 		Code: httpCode,
// 		Msg:  msg,
// 		Data: data,
// 	})
// 	return
// }

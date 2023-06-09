package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/siruspen/logrus"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func BindData(c *gin.Context, req interface{}) bool {
	if c.ContentType() != "application/json" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		err := NewUnsupportedMediaType(msg)

		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return false
	}
	if err := c.Bind(req); err != nil {
		log.Printf("Error binding data: %+v\n", err)
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			var invalidArgs []invalidArgument
			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					fmt.Sprintf("%v", err.Value()),
					err.Tag(),
					err.Param(),
				})
			}
			err := NewBadRequest("Invalid request parameters. See invalidArgs")
			c.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false

		}
		fallbackErr := NewInternal()
		c.JSON(fallbackErr.Status(), gin.H{"error": fallbackErr})
		return false
	}
	return true
}

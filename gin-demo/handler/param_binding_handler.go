package handler

import (
	"github.com/eric-chao/go-micro/gin-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BindingTest(c *gin.Context) {
	var param model.Param
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// param
	var schoolId = param.SchoolId
	var examId = param.ExamId

	var index = schoolId%128 + 1
	var tableName = "r_student_question_" + strconv.Itoa(int(index))

	c.JSON(200, gin.H{
		"schoolId":  schoolId,
		"examId":    examId,
		"tableName": tableName,
	})
}

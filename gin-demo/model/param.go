package model

type Param struct {
	SchoolId int64 `form:"schoolId" binding:"required"`
	ExamId   int64 `form:"examId" binding:"required"`
}

package models

type InputMessage struct{
	Question string `form:"question" binding:"required"`
}

type InputAnswer struct{
	Answer string `form:"answer" binding:"required"`
}


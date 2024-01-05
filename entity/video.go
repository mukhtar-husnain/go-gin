package entity

type Person struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=15"`
	Description string `json:"description" binding:"max=50"`
	Url         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}

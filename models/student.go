package models

type Student struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
	Batch    string `json:"batch"`
}

package models

// Post struct
type Post struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageApp    string `json:"imageapp"`
	ImageWeb    string `json:"imageweb"`
}

// CreatePostInput struct
type CreatePostInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageApp    string `json:"imageapp" binding:"required"`
	ImageWeb    string `json:"imageweb" binding:"required"`
}

//UpdatePostInput struct
type UpdatePostInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageApp    string `json:"imageapp"`
	ImageWeb    string `json:"imageweb"`
}

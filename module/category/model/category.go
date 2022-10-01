package categorymodel

type Category struct {
	Id int `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Description string `json:"description" gorm:"column:description;"`
	Image string `json:"image" gorm:"column:image;"`
}

type CategoryCreate struct {
	Id int `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Description *string `json:"description" gorm:"column:description;"`
	Image *string `json:"image" gorm:"column:image;"`
}

type CategoryUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Description *string `json:"description" gorm:"column:description;"`
	Image *string `json:"image" gorm:"column:image;"`
}

func (Category) TableName() string {return "categories"}
func (CategoryUpdate) TableName() string {return "categories"}
func (CategoryCreate) TableName() string {return "categories"}

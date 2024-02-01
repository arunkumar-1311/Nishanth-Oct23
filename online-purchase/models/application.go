package models

type Roles struct {
	RoleID string `gorm:"column:role_id; uniqueIndex;primaryKey; type:varchar" json:"role_id"`
	Role   string `gorm:"column:role; type:varchar" json:"role"`
}

type OrderStatus struct {
	OrderStatusID string `gorm:"column:order_status_id; uniqueIndex;primaryKey; type:varchar" json:"order_id"`
	Status        string `gorm:"column:status; type:varchar" json:"status"`
}

type Brand struct {
	BrandID string `gorm:"column:brand_id; uniqueIndex; primaryKey; type:varchar" json:"brand_id" validate:"required"`
	Name    string `gorm:"column:name; type:varchar" json:"brand_name" validate:"required"`
	Price   int    `gorm:"column:price; type:int" json:"price" validate:"required"`
	Status  bool   `gorm:"column:status; type:boolean;" json:"status" validate:"required"`
}

type Ram struct {
	RamID  string `gorm:"column:ram_id; uniqueIndex; primaryKey; type:varchar" json:"ram_id" validate:"required"`
	Size   string `gorm:"column:size; type:varchar" json:"ram_size" validate:"required"`
	Price  int    `gorm:"column:price; type:int" json:"price" validate:"required"`
	Status bool   `gorm:"column:status; type:boolean;" json:"status" validate:"required"`
}

type Users struct {
	UserID   string `gorm:"column:user_id; uniqueIndex;primaryKey; type:varchar" json:"user_id" validate:"required"`
	UserName string `gorm:"column:name; uniqueIndex; type:varchar" json:"name" validate:"required"`
	Email    string `gorm:"column:email; uniqueIndex; type:varchar" json:"email" validate:"email,required"`
	Password string `gorm:"column:password; type:varchar" json:"password,omitempty" validate:"required"`
	RolesID  string `gorm:"column:role_id; type :varchar" json:"role_id"`
	Roles    Roles  `gorm:"references:role_id" json:"-"`
}

type Orders struct {
	OrderID       string      `gorm:"column:order_id; uniqueIndex;primaryKey; type:varchar" json:"order_id" validate:"required"`
	FullName      string      `gorm:"column:full_name; type:varchar" json:"full_name" validate:"required"`
	UserID        string      `gorm:"column:user_id; type:varchar" json:"user_id" validate:"required"`
	User          Users       `gorm:"references:UserID;" json:"-" validate:"omitempty,uuid4"`
	PhoneNumber   string      `gorm:"column:phone_number; type:varchar" json:"phone_number" validate:"required"`
	Address       string      `gorm:"column:address; type:varchar" json:"address" validate:"required"`
	BrandID       string      `gorm:"column:brand_id; type:varchar" json:"brand_id" validate:"required"`
	Brand         Brand       `gorm:"references:BrandID;" json:"brand" validate:"omitempty,uuid4"`
	RamID         string      `gorm:"column:ram_id; type:varchar" json:"ram_id" validate:"required"`
	Ram           Ram         `gorm:"references:RamID;" json:"ram" validate:"omitempty,uuid4"`
	DVD           bool        `gorm:"column:dvd_rw; type:boolean;" json:"dvd_rw" validate:"required"`
	Total         int         `gorm:"column:total; type:int" json:"total" validate:"required"`
	OrderStatusID string      `gorm:"column:order_status_id; type:varchar" json:"order_status_id" validate:"required"`
	OrderStatus   OrderStatus `gorm:"references:OrderStatusID;" json:"order_status" validate:"omitempty,uuid4"`
	Active        bool        `gorm:"column:active; type:boolean;" json:"active" validate:"required"`
}

// These type struct's are used to send and recieve data to the client
type Login struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Claims struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	UsersID string `json:"user_id"`
	RolesID string `json:"role"`
}

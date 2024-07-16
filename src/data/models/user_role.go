package models

type UserRole struct {
	BaseModel
	User   User `gorm:"foreign_key:UserId;constraint:on_update:no action;on_delete:no action"`
	Role   Role `gorm:"foreign_key:RoleId;constraint:on_update:no action;on_delete:no action"`
	UserId int
	RoleId int
}

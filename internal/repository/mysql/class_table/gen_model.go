package class_table

import "time"

// ClassTable

//go:generate gormgen -structs ClassTable -input .
type ClassTable struct {
	Id          int32     //
	Name        string    //
	Description string    //
	CreateTime  time.Time `gorm:"time"` //
}

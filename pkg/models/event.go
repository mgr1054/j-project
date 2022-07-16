package models

type Event struct {
	ID			int 		`json:"id"	gorm:"primary_key; auto_increment; not_null"`
	Band_Name	string		`json:"band_name"`
	Location	string		`json:"location"`
	City		string		`json:"city"`
	Capacity	int			`json:"capacity"`
	Date 		string		`json:"date"`
}	
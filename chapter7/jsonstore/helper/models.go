package helper

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	//ShipmentID uint64
	Packages []Package /*`gorm:"foreignKey:ShipmentID"`*/ // Define the one-to-many relationship
	Data     string                                       `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Package struct {
	gorm.Model
	//ShipmentID uint // Foreign key for the Shipment
	//Shipment   *Shipment `gorm:"foreignKey:ShipmentID;belongsTo"` // Reverse relationship
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

//GORM creates tables with plural names
//Use this to suppress it

func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

func InitDB() (*gorm.DB, error) {
	var err error
	dsn := "host=localhost user=**** password=**** dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Shipment{}, &Package{})
	return db, nil
}

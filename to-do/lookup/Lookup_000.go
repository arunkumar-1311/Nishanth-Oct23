package lookup

import "gorm.io/gorm"

func (Empty) Lookup_000(db *gorm.DB) error {
	if err := db.AutoMigrate(&Lookup{}); err != nil {
		return err
	}
	return nil
}

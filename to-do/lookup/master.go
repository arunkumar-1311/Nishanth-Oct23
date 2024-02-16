package lookup

type Empty struct{}

type Lookup struct {
	Name    string `gorm:"column:name"`
	Version string `gorm:"column:version"`
}

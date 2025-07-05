

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Quote{}, &User{})
}
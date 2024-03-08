package schemas

type User struct {
	Id        int
	Email     string `gorm:"unique"`
	Password  string
	Providers []Provider `gorm:"many2many:user_providers;"`
	Files     []File     `gorm:"foreignKey:UserID"`
}

type File struct {
	Id        int
	Name      string
	FileName  string
	CreatedAt string
	UserID    int
}

type Provider struct {
	Id   int
	Name string
}

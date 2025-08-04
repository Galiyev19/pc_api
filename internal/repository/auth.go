package repo


func AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo{
	return &AuthRepo{
		db:db,
	}
}
package data

import "gorm.io/gorm"

var db *gorm.DB

type Store struct {
	Article Article
}

func New(conn *gorm.DB) *Store {
	db = conn
	autoMigrate()
	return &Store{Article: Article{}}
}

func autoMigrate() {
	db.AutoMigrate(Article{})
	migrateScript()
}

func migrateScript() {
	db.Exec(`CREATE EXTENSION pg_trgm;
	ALTER TABLE article ADD text_search tsvector 
		GENERATED ALWAYS AS	(
			setweight(to_tsvector('indonesian', coalesce(title, '')), 'A') || ' ' ||
			setweight(to_tsvector('indonesian', coalesce(slug, '')), 'B') || ' ' || 
			setweight(to_tsvector('indonesian', coalesce(description, '')), 'C') :: tsvector
		) STORED;
	CREATE INDEX idx_text_search ON article USING GIN(text_search);`)
}

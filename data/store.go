package data

import "gorm.io/gorm"

var db *gorm.DB

type Store struct {
	Article Article
	Putusan Putusan
}

func New(conn *gorm.DB) *Store {
	db = conn
	autoMigrate()
	return &Store{Article: Article{}, Putusan: Putusan{}}
}

func autoMigrate() {
	db.AutoMigrate(Article{}, Putusan{})
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
	CREATE INDEX idx_text_search ON article USING GIN(text_search);
	ALTER TABLE putusan ADD text_search tsvector 
		GENERATED ALWAYS AS	(
			setweight(to_tsvector('indonesian', coalesce(title, '')), 'A') || ' ' ||
			setweight(to_tsvector('indonesian', coalesce(slug, '')), 'B') || ' ' || 
			setweight(to_tsvector('indonesian', coalesce(description, '')), 'C') :: tsvector
		) STORED;
	CREATE INDEX idx_text_search ON article USING GIN(text_search); `)
}

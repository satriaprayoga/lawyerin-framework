package data

import "gorm.io/gorm"

var db *gorm.DB

type Store struct {
	Article    Article
	Putusan    Putusan
	Peraturan  Peraturan
	FileUpload FileUpload
	Firm       Firm
}

func New(conn *gorm.DB) *Store {
	db = conn
	autoMigrate()
	return &Store{
		Article:    Article{},
		Putusan:    Putusan{},
		Peraturan:  Peraturan{},
		FileUpload: FileUpload{},
		Firm:       Firm{},
	}
}

func autoMigrate() {
	db.AutoMigrate(Article{},
		Putusan{},
		Peraturan{},
		FileUpload{},
		Firm{})
	migrateScript()
}

func migrateScript() {
	db.Exec(`CREATE EXTENSION pg_trgm;
	CREATE EXTENSION cube; 
	CREATE EXTENSION earthdistance;
	ALTER TABLE article ADD text_search tsvector 
		GENERATED ALWAYS AS	(
			setweight(to_tsvector('indonesian', coalesce(title, '')), 'A') || ' ' ||
			setweight(to_tsvector('indonesian', coalesce(slug, '')), 'B') || ' ' || 
			setweight(to_tsvector('indonesian', coalesce(description, '')), 'C') :: tsvector
		) STORED;
	CREATE INDEX idx_article_text_search ON article USING GIN(text_search);
	ALTER TABLE putusan ADD text_search tsvector 
		GENERATED ALWAYS AS	(
			setweight(to_tsvector('indonesian', coalesce(title, '')), 'A') || ' ' ||
			setweight(to_tsvector('indonesian', coalesce(slug, '')), 'B') || ' ' || 
			setweight(to_tsvector('indonesian', coalesce(description, '')), 'C') :: tsvector
		) STORED;
	CREATE INDEX idx_putusan_text_search ON putusan USING GIN(text_search); 
	ALTER TABLE peraturan ADD text_search tsvector 
		GENERATED ALWAYS AS	(
			setweight(to_tsvector('indonesian', coalesce(title, '')), 'A') || ' ' ||
			setweight(to_tsvector('indonesian', coalesce(slug, '')), 'B') || ' ' || 
			setweight(to_tsvector('indonesian', coalesce(description, '')), 'C') :: tsvector
		) STORED;
	CREATE INDEX idx_peraturan_text_search ON putusan USING GIN(text_search);`)
}

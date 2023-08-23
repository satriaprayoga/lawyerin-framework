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
	db.Exec(`CREATE EXTENSION IF NOT EXISTS pg_trgm;
	CREATE EXTENSION IF NOT EXISTS cube; 
	CREATE EXTENSION IF NOT EXISTS earthdistance;
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
	CREATE INDEX idx_peraturan_text_search ON putusan USING GIN(text_search);
	CREATE OR REPLACE FUNCTION order_by_distance(latitude numeric, longitude numeric)
 RETURNS TABLE(firm_id int8,firm_name varchar, address varchar, city varchar, province varchar, lat numeric, lng numeric, distance numeric)
 LANGUAGE sql
AS $$
SELECT firm_id,firm_name, address, city, province, lat, lng, 6371 * acos(cos(radians(latitude)) * cos(radians(lat)) * cos(radians(lng) - radians(longitude)) + sin(radians(latitude)) * sin(radians(lat))) AS distance
FROM firm
ORDER BY distance ASC;
$$
`)
}

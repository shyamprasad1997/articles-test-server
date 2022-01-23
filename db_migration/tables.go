package main

import (
	"database/sql"
	"fmt"
)

func CreateTables(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS user_app
	(
		user_id integer NOT NULL DEFAULT nextval('user_app_user_id_seq'::regclass),
		name text COLLATE pg_catalog."default" NOT NULL,
		email text COLLATE pg_catalog."default" NOT NULL,
		is_admin boolean,
		CONSTRAINT user_app_pkey PRIMARY KEY (user_id),
		CONSTRAINT email_unique UNIQUE (email)
	)`)
	if err != nil {
		fmt.Println("Failed to create table user_app")
		return err
	}
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS articles
	(
		id_article integer NOT NULL DEFAULT nextval('articles_id_article_seq'::regclass),
		title text COLLATE pg_catalog."default",
		"desc" text COLLATE pg_catalog."default",
		created_at timestamp without time zone,
		created_by integer,
		approval_status boolean DEFAULT false,
		CONSTRAINT articles_pkey PRIMARY KEY (id_article),
		CONSTRAINT created_by_foreign_key FOREIGN KEY (created_by)
			REFERENCES public.user_app (user_id) MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE NO ACTION
			NOT VALID
	)`)
	if err != nil {
		fmt.Println("Failed to create table articles")
		return err
	}
	return nil
}

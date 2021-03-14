package postgres

import (
	"database/sql"
	"github.com/pkg/errors"
	"microblog/types"
)

type PostgreS struct {
	db *sql.DB
}

func NewSQL(psqlInfo string) (*PostgreS, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "err with Open DB")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "err with ping DB")
	}
	return &PostgreS{db}, nil
}

func (p *PostgreS) Close() error {
	return p.db.Close()
}

func (p *PostgreS) Showpost(showPost *types.Blog, vars map[string]string) (*types.Blog, error) {
	res, err := p.db.Query("select * from blog where id = $1", vars["id"])
	if err != nil {
		return nil, errors.Wrap(err, "Err with SQL in Showpost p.db.Query")
	}
	defer res.Close()
	for res.Next() {
		var text types.Blog
		err = res.Scan(&text.Id, &text.Text, &text.Anous, &text.FullText, &text.Now, &text.Username)
		if err != nil {
			return nil, errors.Wrap(err, "Err with SQL in Showpost res.Scan")
		}
		showPost = &text
	}
	return showPost, err
}

func (p *PostgreS) Index(posts []types.Blog) ([]types.Blog, error) {
	res, err := p.db.Query("select * from blog")
	if err != nil {
		return nil, errors.Wrap(err, "Err with SQL in Index p.db.Query")
	}
	defer res.Close()
	for res.Next() {
		var text types.Blog
		err = res.Scan(&text.Id, &text.Text, &text.Anous, &text.FullText, &text.Now, &text.Username)
		if err != nil {
			return nil, errors.Wrap(err, "Err with SQL in Index res.Scan")
		}
		posts = append(posts, text)
	}
	return posts, nil
}

func (p *PostgreS) Saveparamsfromblog(text *types.Blog) error {
	res, err := p.db.Query("insert into blog (title, anous, full_text, datenow, username) values ($1, $2, $3,$4,$5)", text.Text, text.Anous, text.FullText, text.Now, text.Username)
	if err != nil {
		errors.Wrap(err, "Err with SQL in Saveparamsfromblog p.db.Query")
	}
	defer res.Close()
	return err
}

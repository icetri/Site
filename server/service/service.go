package service

import (
	"fmt"
	"github.com/pkg/errors"
	"microblog/postgres"
	"microblog/types"
)

type Service struct {
	p *postgres.PostgreS
}

func NewService(pg *postgres.PostgreS) (*Service, error) {
	return &Service{
		p: pg,
	}, nil
}

func (s *Service) Saveparamsfromblog(text *types.Blog) error {
	if text.Text == "" || text.Anous == "" || text.FullText == "" {
		return fmt.Errorf("Не все данные заполнены")
	} else {
		err := s.p.Saveparamsfromblog(text)
		if err != nil {
			return errors.Wrap(err, "Err with Service in Saveparamsfromblog s.p.Saveparamsfromblog")
		}
	}
	return nil
}

func (s *Service) Showpost(showpost *types.Blog, vars map[string]string) (*types.Blog, error) {
	blog, err := s.p.Showpost(showpost, vars)
	if err != nil {
		return nil, errors.Wrap(err, "Err with Service in Showpost s.p.Showpost")
	}
	return blog, err
}

func (s *Service) Index(post []types.Blog) ([]types.Blog, error) {
	posts, err := s.p.Index(post)
	if err != nil {
		return nil, errors.Wrap(err, "Err with Service Index s.p.Index")
	}
	return posts, err
}

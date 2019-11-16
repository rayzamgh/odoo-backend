package service

import "github.com/odoo-assignment/internal/app/project"

func (s *Service) FetchStoreKeluhan(keluhan *project.Keluhan) (*project.Keluhan, error) {
	return s.Repo.FetchStoreKeluhan(keluhan)
}

package service

import "github.com/odoo-backend/internal/app/project"

func (s *Service) FetchStoreKeluhan(keluhan *project.Keluhan) (*project.Keluhan, error) {
	return s.Repo.FetchStoreKeluhan(keluhan)
}

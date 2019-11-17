package service

import (
	"github.com/odoo-backend/internal/app/project"
)

func (s *Service) FetchShowPertanyaanJawaban(pertanyaan string) (string, error) {
	return s.Repo.FetchShowPertanyaanJawaban(pertanyaan)
}

func (s *Service) FetchStorePertanyaanJawaban(data *project.PertanyaanJawaban) (*project.PertanyaanJawaban, error) {
	return s.Repo.FetchStorePertanyaanJawaban(data)
}

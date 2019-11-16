package service

import "github.com/odoo-assignment/internal/app/project"

func (s *Service) ShowAllPengguna() ([]*project.Pengguna, error) {
	return s.Repo.ShowAllPengguna()
}
func (s *Service) FetchStorePengguna(pengguna *project.Pengguna) (*project.Pengguna, error) {
	return s.Repo.FetchStorePengguna(pengguna)
}

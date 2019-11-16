package project

//
// Service contains services to manage data in the repository.
type Service interface {
	ChatbotService
	Close() error
}

type ChatbotService interface {
	FetchShowPertanyaanJawaban(pertanyaan string) (string, error)
	FetchStorePertanyaanJawaban(*PertanyaanJawaban) (*PertanyaanJawaban, error)
	FetchStoreKeluhan(*Keluhan) (*Keluhan, error)
	ShowAllPengguna() ([]*Pengguna, error)
	FetchStorePengguna(*Pengguna) (*Pengguna, error)
}

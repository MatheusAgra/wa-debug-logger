package data_usecases

import (
	data_protocols "boostify/src/data/protocols"
	domain_usecases "boostify/src/domain/usecases"

	"go.mau.fi/whatsmeow"
)

type dbWhatsappBot struct {
	whatsappConnectionRepository data_protocols.DbWhatsappConnectionRepository
}

func NewDbWhatsappBot(
	whatsappConnectionRepository data_protocols.DbWhatsappConnectionRepository,
) domain_usecases.Bot {
	return &dbWhatsappBot{
		whatsappConnectionRepository: whatsappConnectionRepository,
	}
}

func (w *dbWhatsappBot) Start(numberPhone string, connectionMode string) (client *whatsmeow.Client, err error) {
	return w.whatsappConnectionRepository.Start(numberPhone, connectionMode)
}

func (w *dbWhatsappBot) Stop() error {
	return w.whatsappConnectionRepository.Stop()
}

func (w *dbWhatsappBot) Kill() error {
	return w.whatsappConnectionRepository.Kill()
}

func (w *dbWhatsappBot) Reload() error {
	return w.whatsappConnectionRepository.Reload()
}

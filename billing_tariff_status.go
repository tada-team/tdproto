package tdproto

// Tariff status
type TariffStatus string

const (
	// Tariff is active
	ActiveTariff TariffStatus = "Active"

	// Tariff in archive
	ArchiveTariff TariffStatus = "Archive"
)

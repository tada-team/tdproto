package tdproto

// Tariff status
type TariffStatus string

const (
	// Tariff is active
	ActiveTariff TariffStatus = "TARIFF_STATUS_ACTIVE"

	// Tariff in archive
	ArchiveTariff TariffStatus = "TARIFF_STATUS_ARCHIVE"
)

package rescuetime

import (
	"github.com/bionic-dev/bionic/providers/provider"
	"gorm.io/gorm"
)

const name = "rescuetime"
const tablePrefix = "rescuetime_"

type rescuetime struct {
	provider.Database
}

func New(db *gorm.DB) provider.Provider {
	return &rescuetime{
		Database: provider.NewDatabase(db),
	}
}

func (rescuetime) Name() string {
	return name
}

func (rescuetime) TablePrefix() string {
	return tablePrefix
}

func (rescuetime) ExportDescription() string {
	return "https://www.rescuetime.com/accounts/your-data => \"Your Logged Time\" => \"Activity report history\""
}

func (p *rescuetime) Migrate() error {
	err := p.DB().AutoMigrate(
		&ActivityHistoryItem{},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *rescuetime) ImportFns(inputPath string) ([]provider.ImportFn, error) {
	if provider.IsPathDir(inputPath) {
		return nil, provider.ErrInputPathShouldBeFile
	}

	return []provider.ImportFn{
		provider.NewImportFn(
			"Activity History",
			p.importActivityHistory,
			inputPath,
		),
	}, nil
}

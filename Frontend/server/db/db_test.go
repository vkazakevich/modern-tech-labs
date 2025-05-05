package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vkazakevich/ebiznes/Go/utils"
)

func TestInitDB(t *testing.T) {
	assert.NotPanics(t, func() { 
		db := InitTestDB()
		utils.MigrateDB(db)
	})
}

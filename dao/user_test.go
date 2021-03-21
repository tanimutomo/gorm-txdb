package dao

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDao_Get(t *testing.T) {
	const testDBName = "user_dao_%s"
	tests := []struct {
		name   string
		give   int64
		wantID int64
		err    bool
	}{
		{
			name:   "success",
			give:   int64(1),
			wantID: int64(1),
		},
		{
			name: "not found",
			give: int64(2),
			err:  true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db := prepare(fmt.Sprintf(testDBName, tt.name))
			d := NewUser(db)

			got, aerr := d.Get(tt.give)
			if tt.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
				assert.Equal(t, tt.wantID, got.ID)
			}
		})
	}
}

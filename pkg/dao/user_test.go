package dao

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanimutomo/gorm-txdb/pkg/entity"
	"github.com/tanimutomo/gorm-txdb/pkg/mysql"
)

func TestUserDao_Get(t *testing.T) {
	t.Parallel()

	const testDBName = "user_dao_%s"

	tests := []struct {
		name   string
		data   []interface{}
		give   int64
		wantID int64
		err    bool
	}{
		{
			name: "success",
			data: []interface{}{
				&entity.User{ID: int64(1), Name: "foo"},
			},
			give:   int64(1),
			wantID: int64(1),
		},
		{
			name: "not found",
			give: int64(1),
			err:  true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, err := mysql.NewTest(fmt.Sprintf(testDBName, tt.name))
			if err != nil {
				t.Fatal(err)
			}
			for _, d := range tt.data {
				if err := db.Create(d).Error; err != nil {
					t.Fatal(err)
				}
			}

			got, aerr := NewUser(db).Get(tt.give)
			if tt.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
				assert.Equal(t, tt.wantID, got.ID)
			}
		})
	}
}

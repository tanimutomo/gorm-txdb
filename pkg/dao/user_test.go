package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanimutomo/gorm-txdb/pkg/entity"
)

func TestUserDao_Get(t *testing.T) {
	t.Parallel()

	seeds := []interface{}{
		&entity.User{ID: int64(1), Name: "foo"},
	}
	db, err := prepare("user_dao_get", seeds)
	if err != nil {
		t.Fatal(err)
	}
	d := NewUser(db)

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
			t.Parallel()

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

func TestUserDao_Create(t *testing.T) {
	t.Parallel()

	seeds := []interface{}{
		&entity.User{Name: "exist"},
	}
	db, err := prepare("user_dao_create", seeds)
	if err != nil {
		t.Fatal(err)
	}
	d := NewUser(db)

	tests := []struct {
		name string
		give string
		err  bool
	}{
		{
			name: "success",
			give: "non exist",
		},
		{
			name: "duplicated",
			give: "exist",
			err:  true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			aerr := d.Create(tt.give)
			if tt.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
				assert.NoError(t, db.First(&entity.User{}, "name = ?", tt.give).Error)
			}
		})
	}
}

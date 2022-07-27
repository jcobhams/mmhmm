package database

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jcobhams/mmhmm/models"
)

func TestDb_Write(t *testing.T) {
	db := New()
	type testCase struct {
		key string
		val models.Note
	}
	tests := []testCase{
		{
			key: "key1",
			val: models.Note{
				Id:     "id",
				UserId: "userId",
				Title:  "title",
				Body:   "body",
			},
		},
		{
			key: "key2",
			val: models.Note{
				Id:     "id2",
				UserId: "userId",
				Title:  "title2",
				Body:   "body2",
			},
		},
	}
	for idx, tt := range tests {
		err := db.Write(tt.key, tt.val)
		assert.NoError(t, err)
		assert.Equal(t, idx+1, len(db.store))
	}
}

func TestDb_ReadAndDelete(t *testing.T) {
	db := New()

	type seeder struct {
		key string
		val models.Note
	}
	notes := []seeder{
		{
			key: "key1",
			val: models.Note{
				Id:     "id",
				UserId: "userId",
				Title:  "title",
				Body:   "body",
			},
		},
		{
			key: "key2",
			val: models.Note{
				Id:     "id2",
				UserId: "userId",
				Title:  "title2",
				Body:   "body2",
			},
		},
		{
			key: "key3",
			val: models.Note{
				Id:     "id2",
				UserId: "userIdX",
				Title:  "title2",
				Body:   "body2",
			},
		},
	}
	for _, tt := range notes {
		db.Write(tt.val.UserId, tt.val)
	}

	vals, err := db.Read("userId")
	assert.NoError(t, err)
	assert.Equal(t, len(notes)-1, len(vals))

	db.Delete("userId")
	vals, err = db.Read("userId")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(vals))

}

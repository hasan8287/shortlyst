package saldo

import (
	"context"
	"shortlyst/pkg/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_CheckCoin(t *testing.T) {
// 	repoMoq := &repo.SaldoRepoMock{
// 		GetFunc: func(ctx context.Context, val int) (model.Saldo, error) {
// 			return model.Saldo{Value: val, Count: 2}, nil
// 		},
// 	}

// 	svc := NewSaldoService(repoMoq)

// 	t.Run("Success upsert data", func(t *testing.T) {
// 		_, err := svc.CheckCoin(context.Background(), 300)

// 		assert.NoError(t, err)
// 	})
// }
func Test_Upsert(t *testing.T) {
	repoMoq := &repo.SaldoRepoMock{
		// UpSertFunc: func(ctx context.Context, data model.Saldo) (model.Saldo, error) {
		// 	return data, nil
		// },
	}

	svc := NewSaldoService(repoMoq)

	contentSuccsess := `10|10\n50|20\n500 1`
	contentFailed := `a\nb|20\nc|1`

	t.Run("Success upsert data", func(t *testing.T) {
		data, err := svc.Upsert(context.Background(), contentSuccsess)

		assert.NoError(t, err)
		assert.Equal(t, len(data) > 1, true)
	})

	t.Run("Failed upsert data", func(t *testing.T) {
		data, err := svc.Upsert(context.Background(), contentFailed)

		assert.NoError(t, err)
		assert.Equal(t, len(data) > 1, false)
	})

}

package gateway

import "github.com/kahbum/ms_walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

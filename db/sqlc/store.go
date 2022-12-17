package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxparams containts the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

// TranferTx perform a money ransfer from one account to the other
// It's create a transfer record, add account entries, apdate account's balance with a single db transactions
func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		fmt.Println(txName, " create transfer")
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			Amount:        args.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, " create entry1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.FromAccountID,
			Amount:    -args.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, " create entry2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.ToAccountID,
			Amount:    args.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, " get account 1 for update")
		account1, err := q.GetAccountForUpdate(ctx, args.FromAccountID)

		if err != nil {
			return err
		}

		fmt.Println(txName, " update account 1")
		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			Balance: account1.Balance - args.Amount,
			ID:      args.FromAccountID,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, " get account 2 for update")
		account2, err := q.GetAccountForUpdate(ctx, args.ToAccountID)

		if err != nil {
			return err
		}

		fmt.Println(txName, " update account 2")
		result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			Balance: account2.Balance + args.Amount,
			ID:      args.ToAccountID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

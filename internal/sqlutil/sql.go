// Copyright 2020 The Matrix.org Foundation C.I.C.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlutil

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// ErrUserExists is returned if a username already exists in the database.
var ErrUserExists = errors.New("username already exists")

// A Transaction is something that can be committed or rolledback.
type Transaction interface {
	// Commit the transaction
	Commit() error
	// Rollback the transaction.
	Rollback() error
}

// EndTransaction ends a transaction.
// If the transaction succeeded then it is committed, otherwise it is rolledback.
// You MUST check the error returned from this function to be sure that the transaction
// was applied correctly. For example, 'database is locked' errors in sqlite will happen here.
func EndTransaction(txn Transaction, succeeded *bool) error {
	if *succeeded {
		return txn.Commit()
	} else {
		return txn.Rollback()
	}
}

// EndTransactionWithCheck ends a transaction and overwrites the error pointer if its value was nil.
// If the transaction succeeded then it is committed, otherwise it is rolledback.
// Designed to be used with defer (see EndTransaction otherwise).
func EndTransactionWithCheck(txn Transaction, succeeded *bool, err *error) {
	if e := EndTransaction(txn, succeeded); e != nil && *err == nil {
		*err = e
	}
}

// WithTransaction runs a block of code passing in an SQL transaction
// If the code returns an error or panics then the transactions is rolledback
// Otherwise the transaction is committed.
func WithTransaction(db *sql.DB, fn func(txn *sql.Tx) error) (err error) {
	txn, err := db.Begin()
	if err != nil {
		return fmt.Errorf("sqlutil.WithTransaction.Begin: %w", err)
	}
	succeeded := false
	defer EndTransactionWithCheck(txn, &succeeded, &err)

	err = fn(txn)
	if err != nil {
		return
	}

	succeeded = true
	return
}

// TxStmt wraps an SQL stmt inside an optional transaction.
// If the transaction is nil then it returns the original statement that will
// run outside of a transaction.
// Otherwise returns a copy of the statement that will run inside the transaction.
func TxStmt(transaction *sql.Tx, statement *sql.Stmt) *sql.Stmt {
	if transaction != nil {
		statement = transaction.Stmt(statement)
	}
	return statement
}

// TxStmtContext behaves similarly to TxStmt, with support for also passing context.
func TxStmtContext(context context.Context, transaction *sql.Tx, statement *sql.Stmt) *sql.Stmt {
	if transaction != nil {
		statement = transaction.StmtContext(context, statement)
	}
	return statement
}

// StatementList is a list of SQL statements to prepare and a pointer to where to store the resulting prepared statement.
type StatementList []struct {
	Statement **sql.Stmt
	SQL       string
}

// Prepare the SQL for each statement in the list and assign the result to the prepared statement.
func (s StatementList) Prepare(db *sql.DB) (err error) {
	for _, statement := range s {
		if *statement.Statement, err = db.Prepare(statement.SQL); err != nil {
			err = fmt.Errorf("Error %q while preparing statement: %s", err, statement.SQL)
			return
		}
	}
	return
}

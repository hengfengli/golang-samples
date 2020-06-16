// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spanner

// [START spanner_dml_write_then_read]

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func WriteAndReadUsingDML(w io.Writer, db string) error {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// Insert Record
		stmt := spanner.Statement{
			SQL: `INSERT Singers (SingerId, FirstName, LastName)
				VALUES (11, 'Timothy', 'Campbell')`,
		}
		rowCount, err := txn.Update(ctx, stmt)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%d record(s) inserted.\n", rowCount)

		// Read newly inserted record
		stmt = spanner.Statement{SQL: `SELECT FirstName, LastName FROM Singers WHERE SingerId = 11`}
		iter := txn.Query(ctx, stmt)
		defer iter.Stop()

		for {
			row, err := iter.Next()
			if err == iterator.Done || err != nil {
				break
			}
			var firstName, lastName string
			if err := row.ColumnByName("FirstName", &firstName); err != nil {
				return err
			}
			if err := row.ColumnByName("LastName", &lastName); err != nil {
				return err
			}
			fmt.Fprintf(w, "Found record name with %s, %s", firstName, lastName)
		}
		return err
	})
	return err
}

// [END spanner_dml_write_then_read]

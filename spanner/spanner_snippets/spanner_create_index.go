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

package main

// [START spanner_create_index]

import (
	"context"
	"fmt"
	"io"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func addIndex(w io.Writer, db string) error {
	ctx := context.Background()
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return err
	}

	op, err := adminClient.UpdateDatabaseDdl(ctx, &adminpb.UpdateDatabaseDdlRequest{
		Database: db,
		Statements: []string{
			"CREATE INDEX AlbumsByAlbumTitle ON Albums(AlbumTitle)",
		},
	})
	if err != nil {
		return err
	}
	if err := op.Wait(ctx); err != nil {
		return err
	}
	fmt.Fprintf(w, "Added index\n")
	return nil
}

// [END spanner_create_index]

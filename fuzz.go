// Go TDSQL Driver - A TDSQL-Driver for Go's database/sql package.
//
// Copyright 2020 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

//go:build gofuzz
// +build gofuzz

package tdsql

import (
	"database/sql"
)

func Fuzz(data []byte) int {
	db, err := sql.Open("tdsql", string(data))
	if err != nil {
		return 0
	}
	db.Close()
	return 1
}

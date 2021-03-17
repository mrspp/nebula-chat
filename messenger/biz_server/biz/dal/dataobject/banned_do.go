// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package dataobject

type BannedDO struct {
	Id           int32  `db:"id"`
	Phone        string `db:"phone"`
	BannedTime   int64  `db:"banned_time"`
	Expires      int64  `db:"expires"`
	BannedReason string `db:"banned_reason"`
	Log          string `db:"log"`
	State        int8   `db:"state"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}
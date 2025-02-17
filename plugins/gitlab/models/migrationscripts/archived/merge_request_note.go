/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package archived

import (
	"time"

	"github.com/apache/incubator-devlake/models/migrationscripts/archived"
)

type GitlabMergeRequestNote struct {
	GitlabId        int    `gorm:"primaryKey"`
	MergeRequestId  int    `gorm:"index"`
	MergeRequestIid int    `gorm:"comment:Used in API requests ex. /api/merge_requests/<THIS_IID>"`
	NoteableType    string `gorm:"type:varchar(100)"`
	AuthorUsername  string `gorm:"type:varchar(255)"`
	Body            string
	GitlabCreatedAt time.Time
	Confidential    bool
	Resolvable      bool `gorm:"comment:Is or is not review comment"`
	System          bool `gorm:"comment:Is or is not auto-generated vs. human generated"`

	archived.NoPKModel
}

func (GitlabMergeRequestNote) TableName() string {
	return "_tool_gitlab_merge_request_notes"
}

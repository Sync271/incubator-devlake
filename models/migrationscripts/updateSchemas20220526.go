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

package migrationscripts

import (
	"context"
	"time"

	"github.com/apache/incubator-devlake/models/domainlayer"
	"github.com/apache/incubator-devlake/models/migrationscripts/archived"
	"gorm.io/gorm"
)

type PullRequest20220526 struct {
	domainlayer.DomainEntity
	BaseRepoId     string `gorm:"index"`
	HeadRepoId     string `gorm:"index"`
	Status         string `gorm:"type:varchar(100);comment:open/closed or other"`
	Number         int
	Title          string
	Description    string
	Url            string `gorm:"type:varchar(255)"`
	AuthorName     string `gorm:"type:varchar(100)"`
	AuthorId       string `gorm:"type:varchar(100)"`
	ParentPrId     string `gorm:"index;type:varchar(100)"`
	PullRequestKey int
	CreatedDate    time.Time
	MergedDate     *time.Time
	ClosedDate     *time.Time
	Type           string `gorm:"type:varchar(100)"`
	Component      string `gorm:"type:varchar(100)"`
	MergeCommitSha string `gorm:"type:varchar(40)"`
	HeadRef        string `gorm:"type:varchar(255)"`
	BaseRef        string `gorm:"type:varchar(255)"`
	BaseCommitSha  string `gorm:"type:varchar(40)"`
	HeadCommitSha  string `gorm:"type:varchar(40)"`
}

func (PullRequest20220526) TableName() string {
	return "pull_requests"
}

type updateSchemas20220526 struct{}

func (*updateSchemas20220526) Up(ctx context.Context, db *gorm.DB) error {

	err := db.Migrator().RenameColumn(archived.PullRequest{}, "key", "pull_request_key")
	if err != nil {
		return err
	}
	return nil
}

func (*updateSchemas20220526) Version() uint64 {
	return 20220526000005
}

func (*updateSchemas20220526) Name() string {
	return "update `key` columns to `pull_request_key` at pull_requests"
}

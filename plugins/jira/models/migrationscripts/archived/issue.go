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
	"gorm.io/datatypes"
)

type JiraIssue struct {
	// collected fields
	SourceId                 uint64 `gorm:"primaryKey"`
	IssueId                  uint64 `gorm:"primarykey"`
	ProjectId                uint64
	Self                     string `gorm:"type:varchar(255)"`
	Key                      string `gorm:"type:varchar(255)"`
	Summary                  string
	Type                     string `gorm:"type:varchar(255)"`
	EpicKey                  string `gorm:"type:varchar(255)"`
	StatusName               string `gorm:"type:varchar(255)"`
	StatusKey                string `gorm:"type:varchar(255)"`
	StoryPoint               float64
	OriginalEstimateMinutes  int64  // user input?
	AggregateEstimateMinutes int64  // sum up of all subtasks?
	RemainingEstimateMinutes int64  // could it be negative value?
	CreatorAccountId         string `gorm:"type:varchar(255)"`
	CreatorAccountType       string `gorm:"type:varchar(255)"`
	CreatorDisplayName       string `gorm:"type:varchar(255)"`
	AssigneeAccountId        string `gorm:"type:varchar(255);comment:latest assignee"`
	AssigneeAccountType      string `gorm:"type:varchar(255)"`
	AssigneeDisplayName      string `gorm:"type:varchar(255)"`
	PriorityId               uint64
	PriorityName             string `gorm:"type:varchar(255)"`
	ParentId                 uint64
	ParentKey                string `gorm:"type:varchar(255)"`
	SprintId                 uint64 // latest sprint, issue might cross multiple sprints, would be addressed by #514
	SprintName               string `gorm:"type:varchar(255)"`
	ResolutionDate           *time.Time
	Created                  time.Time
	Updated                  time.Time `gorm:"index"`
	SpentMinutes             int64
	LeadTimeMinutes          uint
	StdStoryPoint            uint
	StdType                  string `gorm:"type:varchar(255)"`
	StdStatus                string `gorm:"type:varchar(255)"`
	AllFields                datatypes.JSONMap

	// internal status tracking
	ChangelogUpdated  *time.Time
	RemotelinkUpdated *time.Time
	archived.NoPKModel
}

type JiraIssueCommit struct {
	archived.NoPKModel
	SourceId  uint64 `gorm:"primaryKey"`
	IssueId   uint64 `gorm:"primaryKey"`
	CommitSha string `gorm:"primaryKey;type:varchar(40)"`
	CommitUrl string `gorm:"type:varchar(255)"`
}

func (JiraIssue) TableName() string {
	return "_tool_jira_issues"
}

func (JiraIssueCommit) TableName() string {
	return "_tool_jira_issue_commits"
}

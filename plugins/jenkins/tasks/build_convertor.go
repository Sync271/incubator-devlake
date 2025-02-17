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

package tasks

import (
	"github.com/apache/incubator-devlake/models/domainlayer"
	"github.com/apache/incubator-devlake/models/domainlayer/devops"
	"github.com/apache/incubator-devlake/models/domainlayer/didgen"
	"github.com/apache/incubator-devlake/plugins/core"
	"github.com/apache/incubator-devlake/plugins/helper"
	models "github.com/apache/incubator-devlake/plugins/jenkins/models"
	"reflect"
)

var ConvertBuildsMeta = core.SubTaskMeta{
	Name:             "convertBuilds",
	EntryPoint:       ConvertBuilds,
	EnabledByDefault: true,
	Description:      "Convert tool layer table jenkins_builds into  domain layer table builds",
}

func ConvertBuilds(taskCtx core.SubTaskContext) error {
	db := taskCtx.GetDb()

	jenkinsBuild := &models.JenkinsBuild{}

	cursor, err := db.Model(jenkinsBuild).Rows()
	if err != nil {
		return err
	}
	defer cursor.Close()

	jobIdGen := didgen.NewDomainIdGenerator(&models.JenkinsJob{})
	buildIdGen := didgen.NewDomainIdGenerator(&models.JenkinsBuild{})

	converter, err := helper.NewDataConverter(helper.DataConverterArgs{
		InputRowType: reflect.TypeOf(models.JenkinsBuild{}),
		Input:        cursor,
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx:   taskCtx,
			Table: RAW_BUILD_TABLE,
		},
		Convert: func(inputRow interface{}) ([]interface{}, error) {
			jenkinsBuild := inputRow.(*models.JenkinsBuild)
			build := &devops.Build{
				DomainEntity: domainlayer.DomainEntity{
					Id: buildIdGen.Generate(jenkinsBuild.JobName, jenkinsBuild.Number),
				},
				JobId:       jobIdGen.Generate(jenkinsBuild.JobName),
				Name:        jenkinsBuild.DisplayName,
				DurationSec: uint64(jenkinsBuild.Duration / 1000),
				Status:      jenkinsBuild.Result,
				StartedDate: jenkinsBuild.StartTime,
				CommitSha:   jenkinsBuild.CommitSha,
			}
			return []interface{}{
				build,
			}, nil
		},
	})
	if err != nil {
		return err
	}

	return converter.Execute()
}

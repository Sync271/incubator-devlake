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

package e2e

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test should only run once main_test is complete and ready

type JenkinsJob struct {
	Id int
}

func TestJenkinsJobs(t *testing.T) {
	var jenkinsJobs []JenkinsJob
	db, err := InitializeDb()
	assert.Nil(t, err)
	if err != nil {
		log.Fatal(err)
	}
	sqlCommand := "SELECT id FROM jenkins_jobs;"
	rows, err := db.Query(sqlCommand)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	assert.Nil(t, err)
	defer rows.Close()
	for rows.Next() {
		var jenkinsJob JenkinsJob
		if err := rows.Scan(&jenkinsJob.Id); err != nil {
			panic(err)
		}
		jenkinsJobs = append(jenkinsJobs, jenkinsJob)
	}
	assert.Equal(t, true, len(jenkinsJobs) > 0)
}

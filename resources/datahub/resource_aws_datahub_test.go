// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	"testing"
)

func TestCheckIfClusterCreationFailed(t *testing.T) {
	type testCase struct {
		description        string
		input              operations.DescribeClusterOK
		shouldContainError bool
	}
	for _, scenario := range []testCase{
		{
			description:        "Test with status: CREATION_FAILED",
			input:              createInputWithStatus("CREATION_FAILED"),
			shouldContainError: true,
		},
		{
			description:        "Test with status: DELETED_ON_PROVIDER",
			input:              createInputWithStatus("DELETED_ON_PROVIDER"),
			shouldContainError: true,
		},
		{
			description:        "Test with status: UPDATE_IN_PROGRESS",
			input:              createInputWithStatus("UPDATE_IN_PROGRESS"),
			shouldContainError: false,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			_, _, err := checkIfClusterCreationFailed(&scenario.input)
			if scenario.shouldContainError && err == nil {
				t.Errorf("Test should contain error but the result does not contain it!")
			}
		})
	}
}

func createInputWithStatus(status string) operations.DescribeClusterOK {
	sum := &models.Cluster{Status: status}
	pl := &models.DescribeClusterResponse{Cluster: sum}
	return operations.DescribeClusterOK{Payload: pl}
}

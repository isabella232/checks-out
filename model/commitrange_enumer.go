// Code generated by "enumer -type=CommitRange"; DO NOT EDIT
/*

SPDX-Copyright: Copyright (c) Capital One Services, LLC
SPDX-License-Identifier: Apache-2.0
Copyright 2017 Capital One Services, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

*/
package model

import (
	"encoding/json"
	"fmt"
)

// CommitRange is an enum.
// It is serialized as a string in JSON and as an integer in SQL.

// CommitRange enum maps.
var (
	strMapCommitRange = map[string]CommitRange{
		"all":  All,
		"head": Head,
	}

	intMapCommitRange = map[CommitRange]string{
		All:  "all",
		Head: "head",
	}
)

// Known says whether or not this value is a known enum value.
func (s CommitRange) Known() bool {
	_, ok := intMapCommitRange[s]
	return ok
}

// String is for the standard stringer interface.
func (s CommitRange) String() string {
	return intMapCommitRange[s]
}

// UnmarshalJSON satisfies the json.Unmarshaler
func (s *CommitRange) UnmarshalJSON(data []byte) error {
	str := ""
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	var ok bool
	*s, ok = strMapCommitRange[str]
	if !ok {
		return fmt.Errorf("Unknown CommitRange enum value: %s", str)
	}
	return nil
}

// MarshalJSON satisfies the json.Marshaler
func (s CommitRange) MarshalJSON() ([]byte, error) {
	if !s.Known() {
		return nil, fmt.Errorf("Unknown CommitRange enum value: %d", int(s))
	}
	name := intMapCommitRange[s]
	return json.Marshal(name)
}

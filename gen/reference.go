/*
 *    Copyright 2019 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package gen

import (
	"github.com/google/gofuzz"
	"github.com/insolar/insolar/core"
)

// ID generates random id.
func ID() (id core.RecordID) {
	fuzz.New().Fuzz(&id)
	return
}

// JetID generates random id.
func JetID() (id core.JetID) {
	fuzz.New().Fuzz(&id)
	return
}

// Reference generates random reference.
func Reference() (ref core.RecordRef) {
	fuzz.New().Fuzz(&ref)
	return
}

/*
Copyright 2017 Andrii Zablodskyi (andrey.zablodskiy@gmail.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchIpLocation(t *testing.T) {
	results := searchIpLocation("8.8.8.8", -1)

	assert.Equal(t, 1, len(results))
	assert.Equal(t, 1, len(results[0].ExtrasOrder))

	extra := results[0].ExtrasOrder[0]
	assert.NotEmpty(t, extra)
	assert.Equal(t, 1, len(results[0].Extras))
	assert.NotEmpty(t, results[0].Extras[extra])
}

func TestIsValidIp4(t *testing.T) {
	golds := []struct {
		ipstring string
		isValid  bool
	}{
		{"9.56.24.100", true},
		{"2001:0db8:0a0b:12f0:0000:0000:0000:0001", false},
		{"not an ip, obviously", false},
	}

	for _, gold := range golds {
		assert.Equal(t, gold.isValid, isValidIp4(gold.ipstring))
	}
}

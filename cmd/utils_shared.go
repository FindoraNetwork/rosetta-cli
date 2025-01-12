// Copyright 2022 Findora, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

func isEmpty(s string) bool {
	return s == ""
}

func isNegative(n int64) bool {
	return n < 0
}

func isEqual(s1 string, s2 string) bool {
	return s1 == s2
}

func isUTXO() bool {
	return Config.CoinSupported
}

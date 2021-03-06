// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/kris-nova/novaarchive/logger"
)

func main() {

	logger.BitwiseLevel = logger.LogEverything
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Debug("Debug logging")

	logger.Writer = os.Stderr
	logger.Critical("Stderr logging")
}

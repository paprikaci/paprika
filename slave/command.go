// Copyright (c) 2014 The AUTHORS
//
// This file is part of paprika.
//
// paprika is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// paprika is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with paprika.  If not, see <http://www.gnu.org/licenses/>.

package slave

import (
	// Stdlib
	"os"
	"runtime"

	// Paprika
	"github.com/paprikaci/paprika/utils"

	// Others
	"github.com/tchap/gocli"
)

var (
	master    string
	token     string
	identity  string
	labels    string
	workspace string
	executors = uint(runtime.NumCPU())
)

var Command = &gocli.Command{
	UsageLine: `
  enslave [-master=URL] [-token=TOKEN] [-identity=IDENTITY] [-labels=LABELS]
          [-workspace=WORKSPACE] [-executors=EXECUTORS]`,
	Short: "run a build slave",
	Long: `
  Start a build slave and connect it to the specified master node.
	`,
	Action: enslaveThisPoorMachine,
}

func init() {
	cmd := Command
	cmd.Flags.StringVar(&master, "master", master, "build master to connect to")
	cmd.Flags.StringVar(&token, "token", token, "build master access token")
	cmd.Flags.StringVar(&identity, "identity", identity, "build slave identity; must be unique")
	cmd.Flags.StringVar(&labels, "labels", labels, "labels to apply to this slave")
	cmd.Flags.StringVar(&workspace, "workspace", workspace, "build workspace")
	cmd.Flags.UintVar(&executors, "executors", executors, "number of jobs that can run in parallel")
}

func enslaveThisPoorMachine(cmd *gocli.Command, args []string) {
	// Make sure there were no arguments specified.
	if len(args) != 0 {
		cmd.Usage()
		os.Exit(2)
	}

	// Read the environment to fill in missing parameters.
	utils.GetenvOrFailNow(&master, "PAPRIKA_MASTER", cmd)
	utils.GetenvOrFailNow(&token, "PAPRIKA_TOKEN", cmd)
	utils.GetenvOrFailNow(&identity, "PAPRIKA_IDENTITY", cmd)
	utils.Getenv(&labels, "PAPRIKA_LABELS")
	utils.GetenvOrFailNow(&workspace, "PAPRIKA_WORKSPACE", cmd)

	// Run the main function.
	enslave()
}
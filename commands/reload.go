// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package commands

//
import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/pagodabox/nanobox-cli/util"
	"github.com/pagodabox/nanobox-golang-stylish"
)

//
var reloadCmd = &cobra.Command{
	Hidden: true,

	Use:   "reload",
	Short: "Reloads the nanobox VM",
	Long: `
Description:
  Reloads the nanobox VM by issuing a "vagrant reload --provision"`,

	PreRun: projectIsCreated,
	Run:    nanoReload,
}

// nanoReload runs 'vagrant reload --provision'
func nanoReload(ccmd *cobra.Command, args []string) {

	// run an init to ensure there is a Vagrantfile
	nanoInit(nil, args)

	fmt.Printf(stylish.Bullet("Reloading nanobox VM"))
	if err := util.RunVagrantCommand(exec.Command("vagrant", "reload", "--provision")); err != nil {
		util.Fatal("[commands/reload] util.RunVagrantCommand() failed", err)
	}
}

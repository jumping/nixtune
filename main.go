/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/cli"
)

var (
	cmdName = "autotune"
	version = "v0.0.0"
)

// Keys for invoking the Lambda scripts. Should only have policies for
// invoking the single lambda script and nothing else.
var (
	awsAPIKey    = ""
	awsSecretKey = ""
	awsRegion    = ""
)

var (
	subscription Subscription
)

func subCmd(cmds ...string) string {
	return fmt.Sprintf("%s %s", cmdName, strings.Join(cmds, " "))
}

func copyright() string {
	return fmt.Sprintf(`Acksin Autotune %s.
Copyright (c) 2015-2016. Acksin.
https://acksin.com/autotune
`, version)
}

func main() {
	subscription = setSubscription(os.Getenv("ACKSIN_FUGUE_API_KEY"))

	c := cli.NewCLI(cmdName, version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"sig": func() (cli.Command, error) {
			return NewSignature(subCmd("signature")), nil
		},
		"list": func() (cli.Command, error) {
			return NewList(subCmd("list")), nil
		},
		// "agent": func() (cli.Command, error) {
		// 	return NewAgent(subCmd("agent")), nil
		// },
	}

	c.HelpFunc = func(commands map[string]cli.CommandFactory) string {
		return fmt.Sprintf("%s\n%s", cli.BasicHelpFunc(cmdName)(commands), copyright())
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)

}

/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

var (
	version = "0.0"
)

func copyright() string {
	return fmt.Sprintf(`Acksin STRUM %s.
Copyright (c) 2016. Acksin.
https://acksin.com/strum
`, version)
}

const (
	statsURL      = "https://api.acksin.com/v1/strum/stats"
	statsDebugURL = "http://localhost:8080/v1/strum/stats"
)

func main() {
	c := cli.NewCLI("strum", version)
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"agent": func() (cli.Command, error) {
			return &agent{}, nil
		},
		"output": func() (cli.Command, error) {
			return &output{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
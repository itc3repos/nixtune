/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package fd

// FD contains the aggregrate file descriptors for the entire system.
type FD struct {
}

// New returns a new FD with the file descriptors.
func New() *FD {
	return &FD{}
}

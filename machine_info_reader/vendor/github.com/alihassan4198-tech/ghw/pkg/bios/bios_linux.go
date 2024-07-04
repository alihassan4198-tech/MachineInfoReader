// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package bios

import "github.com/alihassan4198-tech/ghw/pkg/linuxdmi"

func (i *Info) load() error {
	i.Vendor = linuxdmi.Item(i.ctx, "bios_vendor")
	i.Version = linuxdmi.Item(i.ctx, "bios_version")
	i.Name = linuxdmi.Item(i.ctx, "board_name")
	i.Serialnumber = linuxdmi.Item(i.ctx, "board_serial")
	i.Installdate = linuxdmi.Item(i.ctx, "bios_date")

	return nil
}

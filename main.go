package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/kyleu/libnpn/npnservice-db/authdb"
	"github.com/kyleu/libnpn/npnservice-fs/authfs"
	"github.com/kyleu/libnpn/npnservice/auth"

	"github.com/kyleu/libnpn/npnqueue"

	"github.com/kyleu/libnpn/npnconnection"
	"github.com/kyleu/libnpn/npncontroller"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npndatabase"
	"github.com/kyleu/libnpn/npntemplate/gen/npntemplate"
	"github.com/kyleu/libnpn/npnuser"
	"github.com/kyleu/libnpn/npnweb"
)

func main() {
	_ = npnconnection.Message{}

	npncontroller.InitMime()

	_ = npncore.ErrorDetail{}

	_ = npnuser.Profile{}

	_ = npndatabase.Count{}

	_ = npnqueue.Publisher{}

	b := bytes.Buffer{}
	npntemplate.Boolean(false, "", "", &b)

	var _ auth.Service
	_ = authfs.ServiceFS{}
	_ = authdb.ServiceDatabase{}

	rc := npnweb.RequestContext{
		Title: "Hello!",
	}
	println(fmt.Sprintf("Process: %v - %v", os.Getpid(), rc.Title))
}

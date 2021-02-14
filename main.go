package main

import (
	"bytes"
	"fmt"
	"github.com/kyleu/libnpn/npngraphql"
	"os"

	"github.com/kyleu/libnpn/npnscript/js"
	"github.com/kyleu/libnpn/npnscript/lua"
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

	_ = npngraphql.Service{}

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

	_ = lua.Service{}
	_ = js.Service{}

	err := testbed()
	if err != nil {
		panic(err)
	}

	println(fmt.Sprintf("Process: %v - %v", os.Getpid(), rc.Title))
}

func testbed() error {
	return nil
}

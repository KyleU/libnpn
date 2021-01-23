package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/kyleu/libnpn/npnscript/js"
	"github.com/kyleu/libnpn/npnscript/lua"
	"logur.dev/logur"

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
	logger := logur.NewNoopLogger()

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

	_ = lua.Service{}
	_ = js.Service{}

	err := testbed(logger)
	if err != nil {
		panic(err)
	}

	println(fmt.Sprintf("Process: %v - %v", os.Getpid(), rc.Title))
}

func testbed(logger logur.Logger) error {
	// hangOnASec()
	println("Testbed!")
	ret := "OK"
	println("TESTBED RESULT: " + ret)
	return nil
}

func hangOnASec() {
	println("Press [Enter]")
	_, _ = fmt.Scanln()
}

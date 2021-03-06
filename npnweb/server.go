package npnweb

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/libnpn/npncore"
)

func MakeServer(info AppInfo, r *mux.Router, address string, port uint16) (uint16, error) {
	var msg = "%v is starting in directory [%v] on [%v:%v]"
	if info.Debug() {
		msg += " (verbose)"
	}
	actualPort, l, err := Listen(address, port)
	if err != nil {
		return actualPort, errors.Wrap(err, fmt.Sprintf("unable to listen on port [%v]", actualPort))
	}
	if actualPort != port {
		fmt.Println("port:", actualPort)
	}
	if info != nil {
		dir := info.Files().Root()
		info.Logger().Info(fmt.Sprintf(msg, npncore.AppName, dir, address, actualPort))
	}
	return actualPort, Serve(l, r)
}

func Listen(address string, port uint16) (uint16, net.Listener, error) {
	l, err := net.Listen("tcp", fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		return port, nil, errors.Wrap(err, fmt.Sprintf("unable to listen on port [%v]", port))
	}
	if port == 0 {
		addr := l.Addr().String()
		_, portStr := npncore.SplitStringLast(addr, ':', true)
		actualPort, err := strconv.Atoi(portStr)
		if err != nil {
			return 0, nil, errors.Wrap(err, "invalid port ["+portStr+"]")
		}
		port = uint16(actualPort)
	}
	return port, l, nil
}

func Serve(listener net.Listener, r *mux.Router) error {
	err := http.Serve(listener, r)
	if err != nil {
		return errors.Wrap(err, "unable to run http server")
	}
	return nil
}

package router

import (
	"skynx.io/s-api-go/grpc/network/routing"
	"skynx.io/s-lib/pkg/errors"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/mnet/router/rib"
)

func (r *router) eventProcessor(closeCh <-chan struct{}) {
	for {
		select {
		case nh := <-r.RIB().RelayConnQueue():
			if err := r.relayConnect(nh); err != nil {
				xlog.Warnf("Unable to connect to relay peer: %v", errors.Cause(err))
			}
		case nh := <-r.RIB().RouterConnQueue():
			if err := r.routerConnect(nh); err != nil {
				xlog.Warnf("Unable to connect to router: %v", errors.Cause(err))
			}
		case nh := <-r.RIB().ProxyConnQueue():
			if err := r.proxyConnect(nh); err != nil {
				xlog.Warnf("Unable to connect to iap: %v", errors.Cause(err))
			}
		case evt := <-r.RIB().RouteEventQueue():
			switch evt.Type {
			case rib.RouteEventTypeADD:
				if !r.localForwarding || r.networkInterface == nil {
					continue
				}

				if evt.RouteType == routing.RouteType_STATIC {
					if !r.isValidRouteImport(evt.Addr) {
						continue
					}
				}

				if err := r.networkInterface.routeAdd(cidrIPDst(evt.Addr)); err != nil {
					xlog.Errorf("Unable to add route: %v", errors.Cause(err))
				}
			case rib.RouteEventTypeDELETE:
				if !r.localForwarding || r.networkInterface == nil {
					continue
				}

				if err := r.networkInterface.routeDel(cidrIPDst(evt.Addr)); err != nil {
					xlog.Errorf("Unable to remove route: %v", errors.Cause(err))
				}
			}
		case <-closeCh:
			return
		}
	}
}

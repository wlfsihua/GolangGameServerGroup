package syshandler

import (
	"encoding/json"
	"strings"

	"../../base/ginterface"
	"../../base/handler"
	"../sysdefine"
	"../sysinfo"
	uuid "github.com/satori/go.uuid"
)

// RegisterSubServerHandler handles the request of the registration of subserver
type RegisterSubServerHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *RegisterSubServerHandler) Code() int {
	return sysdefine.RegisterSubServer
}

// OnHandle is called when Handling the packet
func (h *RegisterSubServerHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()
	subServers := h.Node.(ginterface.IGameServer).GetModule(map[uuid.UUID]*sysinfo.SubServerInfo{}).(map[uuid.UUID]*sysinfo.SubServerInfo)

	response := sysdefine.NewRegisterSubServerResultPacket()
	response.Result = sysdefine.Failed

	defer h.Node.(ginterface.IGameServer).SendPacket(peer, response)

	packet := &sysdefine.RegisterSubServerPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("RegisterSubServerHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	if subserverInfo, ok := subServers[peer.GetPeerID()]; ok {
		log.Error("RegisterSubServerHandler.OnHandle(): already registered! peerID = %s, serverType = %d\n", peer.GetPeerID().String(), subserverInfo.ServerType)
		response.Result = sysdefine.RegisterSubServerRegistered
		return false
	}

	subServers[peer.GetPeerID()] = &sysinfo.SubServerInfo{
		SubServerInfoBase: sysinfo.SubServerInfoBase{
			PeerID:     peer.GetPeerID(),
			ServerType: packet.ServerType,
			Address:    peer.GetConn().Request().RemoteAddr[:strings.LastIndex(peer.GetConn().Request().RemoteAddr, ":")],
			Port:       packet.Port,
			ServerName: packet.ServerName,
		},
	}

	response.Result = sysdefine.OK
	return true
}

// NewRegisterSubServerHandler is a constructor of RegisterSubServerHandler
func NewRegisterSubServerHandler(node ginterface.INode) *RegisterSubServerHandler {
	ret := &RegisterSubServerHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}

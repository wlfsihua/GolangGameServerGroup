package syshandler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../sysdefine"
)

// RegisterSubServerResultHandler handles the result of the registration to the master server
type RegisterSubServerResultHandler struct {
	*handler.GameHandler
}

// Code is the associated packet command
func (h *RegisterSubServerResultHandler) Code() int {
	return sysdefine.RegisterSubServerResult
}

// OnHandle is called when Handling the packet
func (h *RegisterSubServerResultHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.(ginterface.IGameServer).GetLogger()

	packet := &sysdefine.RegisterSubServerResultPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("RegisterSubServerResultHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	if packet.Result != sysdefine.OK {
		log.Error("RegisterSubServerResultHandler.OnHandle(): result is not ok! result = %d\n", packet.Result)
		h.Node.(ginterface.IGameServer).Stop()
	}
	return true
}

// NewRegisterSubServerResultHandler is a constructor of RegisterSubServerResultHandler
func NewRegisterSubServerResultHandler(node ginterface.INode) *RegisterSubServerResultHandler {
	ret := &RegisterSubServerResultHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}

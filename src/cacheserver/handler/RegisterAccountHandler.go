package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysdefine"
)

// RegisterAccountHandler handles the registration of the player's account
type RegisterAccountHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *RegisterAccountHandler) Code() int {
	return gamedefine.RegisterAccount
}

// OnHandle is called when Handle()
func (h *RegisterAccountHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Server.GetLogger()

	response := gamedefine.NewRegisterAccountResultPacket()

	packet := &gamedefine.RegisterAccountPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("RegisterAccountHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}
	log.Debug("RegisterAccountHandler.OnHandle(): code = %d, account = %s, password=%s\n", h.Code(), packet.Account, packet.Password)

	response.PeerID = packet.PeerID
	response.Result = sysdefine.OK
	h.Server.SendPacket(peer, response)
	return true
}

// NewRegisterAccountHandler is a constructor of RegisterAccountHandler
func NewRegisterAccountHandler(server ginterface.IGameServer) *RegisterAccountHandler {
	ret := &RegisterAccountHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, server)
	return ret
}

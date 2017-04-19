package bridge

import (
	"github.com/craftsman-li/ggbot/bridge/arg"
	"github.com/craftsman-li/ggbot/bridge/result"
	"github.com/craftsman-li/wechat"
)

//login client will login this uuid
func (w *Wrapper) login(uuid string) *result.Result {

	a := arg.NewArg(arg.Login).Append(`uuid`, uuid)

	r := w.Call(a)

	return r
}

// OpenRedPacket designed to notify wechat app open read packet.
func (w *Wrapper) OpenRedPacket() *result.Result {

	a := arg.NewArg(arg.RedPacket)

	r := w.Call(a)

	return r
}

// SendMessage ...
func (w *Wrapper) SendMessage(msg wechat.Msg) *result.Result {
	return result.NewResultWithError(`unimplement`)
}

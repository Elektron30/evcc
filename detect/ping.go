package detect

import (
	"time"

	"github.com/andig/evcc/util"
	"github.com/go-ping/ping"
)

func init() {
	registry.Add("ping", PingHandlerFactory)
}

func PingHandlerFactory(conf map[string]interface{}) (TaskHandler, error) {
	handler := PingHandler{
		Count:   1,
		Timeout: timeout,
	}

	err := util.DecodeOther(conf, &handler)

	return &handler, err
}

type PingHandler struct {
	Count   int
	Timeout time.Duration
}

func (h *PingHandler) Test(log *util.Logger, ip string) (res []interface{}) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}

	pinger.Count = h.Count
	pinger.Timeout = h.Timeout

	if err = pinger.Run(); err != nil {
		log.FATAL.Println("ping:", err)
		log.FATAL.Println("")
		log.FATAL.Println("In order to run evcc in discovery mode, make sure to allow ping:")
		log.FATAL.Println("")
		log.FATAL.Println("	sudo sysctl -w net.ipv4.ping_group_range=\"0 2147483647\"")
		log.FATAL.Fatalln("")
	}

	stat := pinger.Statistics()

	if stat.PacketsRecv == 0 {
		return nil
	}

	return []interface{}{nil}
}

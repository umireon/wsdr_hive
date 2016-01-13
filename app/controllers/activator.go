package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/umireon/wsdr_hive/app/activator"
)

type Activator struct {
	*revel.Controller
}

func (c Activator) Activate() revel.Result {
	return c.Render()
}

func (c Activator) ActivatePost() revel.Result {
	var config struct {
		System struct {
			Driver string
			Target struct {
				Tx string
				Rx string
			}
		}
		Mod struct {
			Scheme string
		}
		Filter struct {
			SymbolDuration int
			Span           int
			Type           struct {
				Tx string
				Rx string
			}
			Rolloff float64
		}
	}
	c.Params.Bind(&config.System, "System")
	c.Params.Bind(&config.Mod, "Mod")
	c.Params.Bind(&config.Filter, "Filter")
	str, _ := json.Marshal(config)
	activator.Activate(string(str))
	return c.RenderJson("")
}

func (c Activator) Join() revel.Result {
	revel.TRACE.Println("Activate.Join")
	e := <-activator.Join()
	return c.RenderText(e.Message)
}

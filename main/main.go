package main

import (
	t "github.com/TyphoonMC/TyphoonCore"
)

func main() {
	core := t.Init()
	core.SetBrand("exampleserver")

	core.On(func(e *t.PlayerJoinEvent) {
		msg := t.ChatMessage("Welcome ")
		msg.SetExtra([]t.IChatComponent{
			t.ChatMessage(e.Player.GetName()),
			t.ChatMessage(" !"),
		})
		e.Player.SendMessage(msg)
	})

	core.Start()
}

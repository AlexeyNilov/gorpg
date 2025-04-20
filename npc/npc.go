package npc

type NPC struct {
	Name string
	Description string
	Log []string
}

func (n *NPC) Describe() string {
	return n.Description
}

func (n *NPC) LogEvent(event string) {
	n.Log = append(n.Log, event)
}
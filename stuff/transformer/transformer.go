package transformer

type Transformer struct{
	Consumed int
	Produced int
}

func (t *Transformer) Consume() {
	t.Consumed++
}

func (t *Transformer) Produce() {
	t.Produced++
}

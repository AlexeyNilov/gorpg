package transformer

import "github.com/AlexeyNilov/gorpg/stuff/resource"

type Transformer struct {
	Consumed int
	Produced int
}

func (t *Transformer) Consume() {
	t.Consumed++
}

func (t *Transformer) Produce() {
	t.Produced++
}

type ResourceTransformer struct {
	ConsumedCount  int
	Capacity int
	ProducedCount  int
}

func (t *ResourceTransformer) Consume(getter resource.Getter, amount int) {
	t.Capacity += getter.Get(amount)
	t.ConsumedCount++
}

func (t *ResourceTransformer) Produce(putter resource.Putter, amount int) {
	t.Capacity -= putter.Put(amount)
	t.ProducedCount++
}

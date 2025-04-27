package transformer

import "github.com/AlexeyNilov/gorpg/stuff/resource"

type Transformer struct {
	ConsumedCount int
	ProducedCount int
}

func (t *Transformer) Consume() {
	t.ConsumedCount++
}

func (t *Transformer) Produce() {
	t.ProducedCount++
}

type ResourceTransformer struct {
	Transformer
	resource.Resource
}

func (t *ResourceTransformer) Consume(getter resource.Getter, amount int) {
	t.Resource.Put(getter.Get(amount))
	t.ConsumedCount++
}

func (t *ResourceTransformer) Produce(putter resource.Putter, amount int) {
	putter.Put(t.Resource.Get(amount))
	t.ProducedCount++
}

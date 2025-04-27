package resource

type Resource struct{
	Value int
	MaxValue int
}

type Getter interface {
	Get(amount int) int
}

type Putter interface {
	Put(amount int) int
}

func (r *Resource) Get(amount int) int {
	if amount <= 0 {
		return 0
	}
	if amount > r.Value {
		amount = r.Value
	}
	r.Value -= amount
	return amount
}

func (r *Resource) Put(amount int) int {
	if amount <= 0 {
		return 0
	}
	if amount + r.Value > r.MaxValue {
		amount = r.MaxValue - r.Value
	}
	r.Value += amount
	return amount
}
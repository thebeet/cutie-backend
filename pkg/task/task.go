package task

type Component interface {
	Get(id int) (*Task, error)
}

type component struct {
}

func New() Component {
	return &component{}
}

func (c *component) Get(id int) (*Task, error) {
	return nil, nil
}

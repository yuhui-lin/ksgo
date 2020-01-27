package ksgo

type Topology struct {
}

func (t *Topology) AddSource(name string) *Topology {
	return t
}

func (t *Topology) AddSink(name string) *Topology {
	return t
}

func (t *Topology) AddProcessor(name string) *Topology {
	return t
}

func (t *Topology) AddStateStore(name string) *Topology {
	return t
}

func (t *Topology) AddGlobalStore(name string) *Topology {
	return t
}

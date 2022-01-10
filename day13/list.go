package main

type Assets struct {
	assets []Asset
}

func (a *Assets) DoStartWork() {
	for _, item := range a.assets {
		if door, ok := item.(Door); ok {
			door.Open()
		}
	}
}

func (a *Assets) DoEndWork() {
	for _, item := range a.assets {
		if door, ok := item.(Door); ok {
			door.Close()
		}
	}
}

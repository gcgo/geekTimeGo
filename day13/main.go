package main

func main() {
	ss := Assets{
		assets: []Asset{
			DoorGlass{},
			WoodGlass{},
		},
	}
	ss.DoStartWork()
	ss.DoEndWork()
}

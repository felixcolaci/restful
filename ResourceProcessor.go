package restful

type ResourceProcessor interface {
	Process(models []struct{}) ([]RestResource, error)
}

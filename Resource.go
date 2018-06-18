package restful

type RestResource interface {
	Add(link Link)
	EmbedResource(resource RestResource, resourceName string)
	EmbedResources(resources []RestResource, resourceName string)
}


type ResourceCollection struct {
	Page Page `json:"_page,omitempty"`
	Resource
}

type Resource struct {
	Links map[string]Link `json:"_links,omitempty"`
	EmbeddedResources map[string]interface{} `json:"_embedded,omitempty"`
	Interactions map[string]Link `json:"_Interactions,omitempty"`
}

func (r *Resource) Add(link Link) {
	if len(link.Method) > 0 {
		r.addInteraction(link)
	} else {
		r.addLink(link)
	}
}

func (r *Resource) EmbedResource(resource RestResource, resourceName string) {
	if len(r.EmbeddedResources) < 1 {
		r.EmbeddedResources = make(map[string]interface{})
	}
	r.EmbeddedResources[resourceName] = resource
}
func (r *Resource) EmbedResources(resources []RestResource, resourceName string) {
	if len(r.EmbeddedResources) < 1 {
		r.EmbeddedResources = make(map[string]interface{})
	}
	r.EmbeddedResources[resourceName] = resources
}

func (r *Resource) addLink(link Link) {
	if len(r.Links) < 1 {
		r.Links = map[string]Link{link.Rel: link}
	}
	r.Links[link.Rel] = link
}

func (r *Resource) addInteraction(link Link) {
	if len(r.Interactions) < 1 {
		r.Interactions = map[string]Link{link.Rel: link}
	}
	r.Interactions[link.Rel] = link
}
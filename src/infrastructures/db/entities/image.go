package entities

type Image struct {
	Name      string `json:"name" dynamo:"name"`
	URL       string `json:"url" dynamo:"url"`
	CreatedAt string `json:"createdAt" dynamo:"createdAt"`
}

type Images []Image

func (imgs Images) Len() int {
	return len(imgs)
}

func (imgs Images) Swap(i, j int) {
	imgs[i], imgs[j] = imgs[j], imgs[i]
}

func (imgs Images) Less(i, j int) bool {
	return imgs[i].CreatedAt > imgs[j].CreatedAt
}

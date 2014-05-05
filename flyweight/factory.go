package flyweight

type WebsitFactory struct {
	maps map[Websit]Websit
}

func NewWebsitFactory() WebsitFactory {
	maps := make(map[Websit]Websit)
	return WebsitFactory{maps: maps}
}

func (this *WebsitFactory) GetWebsitDescription(websit Websit) {
	maps[i]
}

package playwright

import "fmt"

type Browser struct {
	ChannelOwner
	IsConnected bool
	Contexts    []*BrowserContext
}

func (b *Browser) NewContext() (*BrowserContext, error) {
	channelOwner, err := b.channel.Send("newContext", nil)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %v", err)
	}
	context := channelOwner.(*Channel).object.(*BrowserContext)
	b.Contexts = append(b.Contexts, context)
	return context, nil
}

func (b *Browser) Close() error {
	_, err := b.channel.Send("close", nil)
	return err
}

func (b *Browser) Version() string {
	return b.initializer["version"].(string)
}

func newBrowser(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Browser {
	bt := &Browser{
		IsConnected: true,
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}

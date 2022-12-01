package video

import "net/url"

// AddAuthorName chain function to add author name to an existing video
func (v Video) AddAuthorName(authorName string) Video {
	v.setAuthorName(authorName)
	return v
}

// RemoveAuthorName remove add author name from video
func (v Video) RemoveAuthorName() Video {
	v.removeAuthorName()
	return v
}

// AddProviderName chain function to add provider name to an existing video
func (v Video) AddProviderName(providerName string) Video {
	v.setProviderName(providerName)
	return v
}

// RemoveProviderName remove add provider name from video
func (v Video) RemoveProviderName() Video {
	v.removeProviderName()
	return v
}

// AddDescription chain function to add description to an existing video
func (v Video) AddDescription(description string) Video {
	v.setDescription(description)
	return v
}

// RemoveDescription remove add description from video
func (v Video) RemoveDescription() Video {
	v.removeDescription()
	return v
}

// AddProviderIconUrl chain function to add provider icon url to an existing video
func (v Video) AddProviderIconUrl(providerIconUrl *url.URL) Video {
	v.setProviderIconUrl(providerIconUrl)
	return v
}

// RemoveProviderIconUrl remove add provider icon url from video
func (v Video) RemoveProviderIconUrl() Video {
	v.removeProviderIconUrl()
	return v
}

// AddTitleUrl chain function to add title url to an existing video
func (v Video) AddTitleUrl(titleUrl *url.URL) Video {
	v.setTitleUrl(titleUrl)
	return v
}

// RemoveTitleUrl remove add title url from video
func (v Video) RemoveTitleUrl() Video {
	v.removeTitleUrl()
	return v
}

// AddBlockId chain function to add block id to an existing video
func (v Video) AddBlockId(blockId string) Video {
	v.setBlockId(blockId)
	return v
}

// RemoveBlockId remove add block id from video
func (v Video) RemoveBlockId() Video {
	v.removeBlockId()
	return v
}

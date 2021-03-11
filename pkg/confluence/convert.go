package client

/* import (
	"html"
	"log"
)
 */
/* //ConvertToStorage returns the supplied text in storage format
func (c *ConfluenceClient) ConvertToStorage(body, sourceFormat, title, spaceKey string) (*string, error) {
	switch sourceFormat {
	case "markdown":
		return c.convertFromMarkdown(body, title, spaceKey)

	case "wiki", "storage":
		return c.convertDefault(body, sourceFormat)

	default:
		log.Println("Warning: Format unknown, service may not be able to convert format", sourceFormat)
		return c.convertDefault(body, sourceFormat)
	}
}

func (c *ConfluenceClient) convertDefault(body, sourceFormat string) (*string, error) {
	convertRequest := &ConfluenceConvert{
		Value:          body,
		Representation: sourceFormat,
	}
	_, err := c.doRequest("POST", "/rest/api/contentbody/convert/storage", convertRequest, convertRequest)
	if err != nil {
		return nil, err
	}
	return &convertRequest.Value, nil
}

func (c *ConfluenceClient) convertFromMarkdown(body, title, spaceKey string) (*string, error) {
	log.Println("Warning: This request relies on an undocumented API and is subject to change")
	entityID := c.SearchPages(title, spaceKey)
	convertRequest := &TinyMceRequest{
		EntityID: entityID.Results[0].ID,
		SpaceKey: spaceKey,
		Wiki:     body,
	}

	response, err := c.doRequest("POST", "/rest/tinymce/1/markdownxhtmlconverter", convertRequest, nil)
	if err != nil {
		return nil, err
	}
	result := html.UnescapeString(string(response))
	return &result, nil
}
 */

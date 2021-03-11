package client

import (
    "net/http"
    "net/url"
    "time"

    log "github.com/sirupsen/logrus"
)

//ConfluenceConfig holds the current client configuration
type ConfluenceConfig struct {
    Username string
    Password string
    URL      string
    Debug    bool
}


//ConfluenceClient is the primary client to the Confluence API
type ConfluenceClient struct {
    config  *ConfluenceConfig
    client  *http.Client
}

//Client returns a new instance of the client
func Client(config *ConfluenceConfig) *ConfluenceClient {
    log.SetLevel(log.InfoLevel)
    if config.Debug {
        log.SetLevel(log.DebugLevel)
    }
    return &ConfluenceClient{
        config: config,
        client: &http.Client{
            Timeout: 60 * time.Second,
        },
    }
}

//AddPage adds a new page to the space with the given title
func (c *ConfluenceClient) AddPage(title, spaceKey, body string, ancestor int64) {
    page := newPage(title, spaceKey)
    if ancestor > 0 {
        page.Ancestors = []ConfluencePageAncestor{
            ConfluencePageAncestor{ancestor},
        }
    }
    response := &ConfluencePage{}
    page.Body.Storage.Value = body
    //page.Body.Storage.Representation = "wiki"
    c.doRequest("POST", "/rest/api/content/", page, response)
    log.Println("Confluence page added with ID", response.ID, "and version", response.Version.Number)
}

//UpdatePage adds a new page to the space with the given title
func (c *ConfluenceClient) UpdatePage(title, spaceKey, body string, ID string, version, ancestor int64) {
    page := newPage(title, spaceKey)
    page.ID = ID
    page.Version = &ConfluencePageVersion{version}
    if ancestor > 0 {
        page.Ancestors = []ConfluencePageAncestor{
            ConfluencePageAncestor{ancestor},
        }
    }
    response := &ConfluencePage{}
    page.Body.Storage.Value = body
    //page.Body.Storage.Representation = "wiki"
    c.doRequest("PUT", "/rest/api/content/"+ID, page, response)
    log.Println("Confluence page updated with ID", response.ID, "and version", response.Version.Number)
}

//SearchPages searches for pages in the space that meet the specified criteria
func (c *ConfluenceClient) GetPage(id, expand string) (results *ConfluencePage) {
    results = &ConfluencePage{}
    c.doRequest("GET", "/rest/api/content/"+url.QueryEscape(id)+"?expand="+url.QueryEscape(expand), nil, results)
    return results
}

//CQLSearchPagesBy searches for pages in the space that meet the specified criteria
func (c *ConfluenceClient) SearchPagesByCQL(cql, expand string) (results *ConfluencePagesSearch) {
    results = &ConfluencePagesSearch{}
    c.doRequest("GET", "/rest/api/search?limit=5&cql="+url.QueryEscape(cql)+"&expand=" + url.QueryEscape(expand), nil, &ConfluencePagesSearch{})
    return results
}

//AddAttachment adds an attachment to an existing page
func (c *ConfluenceClient) AddAttachment(content, pageID, filename string) {
    results := &ConfluencePageSearch{}
    c.uploadFile("PUT", "/rest/api/content/"+pageID+"/child/attachment", content, filename, &results)
}

//GetLabel searches for pages in the space that meet the specified criteria
func (c *ConfluenceClient) GetLabel(label string) (results *Label) {
    c.doRequest("GET", "/rest/api/label?name="+label, nil, &Label{})
    return results
}

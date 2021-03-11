package client

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "mime/multipart"
    "net/http"

    log "github.com/sirupsen/logrus"
)

func (c *ConfluenceClient) doRequest(method, url string, content, responseContainer interface{}) ([]byte, error) {
    b := new(bytes.Buffer)
    if content != nil {
        json.NewEncoder(b).Encode(content)
    }
    furl := c.config.URL + url
    log.Println("Full URL", furl)
    log.Println("JSON Content:", b.String())

    request, err := http.NewRequest(method, furl, b)
    request.SetBasicAuth(c.config.Username, c.config.Password)
    request.Header.Add("Content-Type", "application/json; charset=utf-8")
    if err != nil {
        log.Errorln(err)
        return nil, err
    }

    log.Println("Sending request to services...")
    response, err := c.client.Do(request)
    if err != nil {
        log.Errorln(err)
        return nil, err
    }

    defer response.Body.Close()
    log.Println("Response received, processing response...")
    log.Println("Response status code is", response.StatusCode)
    log.Println(response.Status)

    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Errorln(err)
        return nil, err
    }
    if c.config.Debug {
        var prettyJSON bytes.Buffer
        err := json.Indent(&prettyJSON, contents, "", "\t")
        if err != nil {
            log.Errorf("JSON parse error: %e", err)
            return nil, err
        }

        log.Println("Response from service...", string(prettyJSON.Bytes()))
    }

    if response.StatusCode != 200 {
        log.Errorf("Bad response code received from server: %s", response.Status)
        return nil, err
    }
    err = json.Unmarshal(contents, responseContainer)
	if err != nil {
        log.Errorln(err)
        return nil, err
    }
    return contents, nil
}

func (c *ConfluenceClient) uploadFile(method, url, content, filename string, responseContainer interface{}) ([]byte, error) {
    b := new(bytes.Buffer)
    writer := multipart.NewWriter(b)
    part, err := writer.CreateFormFile("file", filename)
    if err != nil {
        log.Errorln(err)
        return nil, err
    }
    part.Write([]byte(content))
    writer.WriteField("minorEdit", "true")
    //writer.WriteField("comment", "test")
    writer.Close()

    furl := c.config.URL + url
    log.Println("Full URL", furl)

    request, err := http.NewRequest(method, furl, b)
    request.SetBasicAuth(c.config.Username, c.config.Password)
    request.Header.Add("Content-Type", writer.FormDataContentType())
    request.Header.Add("X-Atlassian-Token", "nocheck")
    if err != nil {
        log.Errorln(err)
        return nil, err
    }
    log.Println("Sending request to services...")

    response, err := c.client.Do(request)
    if err != nil {
        log.Errorln(err)
        return nil, err
    }
    defer response.Body.Close()
    log.Println("Response received, processing response...")
    log.Println("Response status code is", response.StatusCode)

    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Errorln(err)
        return nil, err
    }
    if response.StatusCode != 200 {
        log.Errorf("Bad response code received from server: %s", response.Status)
        return nil, err
    }
    err = json.Unmarshal(contents, responseContainer)
	if err != nil {
        log.Errorln(err)
        return nil, err
    }
    return contents, nil
}

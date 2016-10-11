package reksai 

import (
    "strings"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "github.com/michaelleah/env"
    "errors"
)

type Summoner struct {
    ProfileIconID int    `json:"profileIconId"`
    SummonerLevel int    `json:"summonerLevel"`
    ID            int64  `json:"id"`
    Name          string `json:"name"`
    RevisionDate  int64  `json:"revisionDate"`
}

type Mastery struct {
    ID   int `json:"id"`
    Rank int `json:"rank"`
}

type MasteryPage struct {
    ID        int     `json:"id"`
    Name      string  `json:"name"`
    Current   bool    `json:"current"`
    Masteries []Mastery `json:"masteries"`
}

type MasteryBook struct {
    SummonerID int           `json:"summonerId"`
    Pages      []MasteryPage `json:"pages"`
}

/**
 * SummonersByName
 * 
 * Method takes one to many summoner names and a region. It then pulls out a list
 * of summoner details for the passed in summoner names and returns them 
 * as a map. Method will return an error if it fails at any point.
 * 
 * @param region string     
 * @param names  ...string
 * 
 * @return summoners map[string]Summoner, err error 
 */
func SummonersByName(region string, names ...string) (summoners map[string]Summoner, err error) {
    API_KEY, err := env.Get("API_KEY")
    if err != nil {
        return nil, errors.New("API Key not configured in .env file.")
    }

    summonerList := strings.Join(names, ",")

    // Build the URL
    url := "https://" + region + ".api.pvp.net/api/lol/"  + region + "/v1.4/summoner/by-name/" + summonerList + "?api_key=" + API_KEY

    // Fire off the Request
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    // Parse the request
    defer resp.Body.Close()
    resp_body, _ := ioutil.ReadAll(resp.Body)

    if err := json.Unmarshal([]byte(resp_body), &summoners); err != nil {
        return nil, err
    }

    return summoners, nil
}

/**
 * SummonersByID
 * 
 * Method takes one to many summoner ids and a region. It then pulls out a list
 * of summoner details for the passed in summoner ids and returns them as a 
 * map. Method will return an error if it fails at any point.
 * 
 * @param region string     
 * @param ids    ...string
 * 
 * @return summoners map[string]Summoner, err error
 */
func SummonersByID(region string, ids ...string) (summoners map[string]Summoner, err error) {
    API_KEY, err := env.Get("API_KEY")
    if err != nil {
        return nil, errors.New("API Key not configured in .env file.")
    }

    IDList := strings.Join(ids, ",")

    // Build the URL
    url := "https://" + region + ".api.pvp.net/api/lol/"  + region + "/v1.4/summoner/" + IDList + "?api_key=" + API_KEY

    // Fire off the Request
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    // Parse the request
    defer resp.Body.Close()
    resp_body, _ := ioutil.ReadAll(resp.Body)

    if err := json.Unmarshal([]byte(resp_body), &summoners); err != nil {
        return nil, err
    }

    return summoners, nil
}

/**
 * MasteriesByID
 * 
 * Method takes one to many summoner ids and a region. It then pulls out a list
 * of mastery pages for the passed in summoner ids and returns them as a map. 
 * Method will return an error if it fails at any point.
 * 
 * @param region string     
 * @param ids    ...string
 * 
 * @return summoners map[string]MasteryBook, err error
 */
func MasteriesByID(region string, ids ...string) (summoners map[string]MasteryBook, err error) {
    API_KEY, err := env.Get("API_KEY")
    if err != nil {
        return nil, errors.New("API Key not configured in .env file.")
    }

    IDList := strings.Join(ids, ",")

    // Build the URL
    url := "https://" + region + ".api.pvp.net/api/lol/" + region + "/v1.4/summoner/" + IDList +"/masteries?api_key=" + API_KEY
    
    // Fire off the Request
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

     defer resp.Body.Close()
    resp_body, _ := ioutil.ReadAll(resp.Body)

    if err := json.Unmarshal([]byte(resp_body), &summoners); err != nil {
        return nil, err
    }

    return summoners, nil
}
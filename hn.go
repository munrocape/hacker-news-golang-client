package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseUrl string
	UserSuffix string
	ItemSuffix string
	MaxSuffix string
	TopSuffix string
	NewSuffix string
	JobSuffix string
	AskSuffix string
	ShowSuffix string
	UpdateSuffix string
}

func NewClient() *Client {
	var c Client
	c.BaseUrl = "https://hacker-news.firebaseio.com/v0/"
	c.UserSuffix = "user/%s.json"
	c.ItemSuffix = "item/%d.json"
	c.MaxSuffix = "maxitem.json"
	c.TopSuffix = "topstories.json"
	c.NewSuffix = "newstories.json"
	c.JobSuffix = "jobstories.json"
	c.AskSuffix = "askstories.json"
	c.ShowSuffix = "showstories.json"
	c.UpdateSuffix = "updates.json"
	return &c
}

func (c *Client) GetResource(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		return contents, err
	}
}

func (c *Client) GetItem(id int) (Item, error) {
	url := c.BaseUrl + fmt.Sprintf(c.ItemSuffix, id)
	rep, err := c.GetResource(url)

	var i Item
	if err != nil {
		return i, err
	}

	err = json.Unmarshal(rep, &i)
	return i, err
}

func (c *Client) GetUser(username string) (User, error) {
	url := c.BaseUrl + fmt.Sprintf(c.UserSuffix, username)
	rep, err := c.GetResource(url)

	var user User
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(rep, &user)
	return user, err
}

// GetTopStories takes an int number and returns an array of up to number ints that represent the current top stories.
// Constraints: 0 <= number <= 500
func (c *Client) GetTopStories(number int) ([]int, error) {
	var top500 []int
	if number > 500 {
		return top500, fmt.Errorf("Number %d greater than maximum 500 items allowed", number)
	}
	
	url := c.BaseUrl + c.TopSuffix
	rep, err := c.GetResource(url)
	
	err = json.Unmarshal(rep, &top500)
	
	if err != nil {
		return nil, err
	}

	return top500[:number], nil
}

// GetNewStories takes an int number and returns an array of up to number ints that represent the newest stories.
// Constraints: 0 <= number <= 500
func (c *Client) GetNewStories(number int) ([]int, error) {
	var top500 []int
	if number > 500 {
		return top500, fmt.Errorf("Number %d greater than maximum 500 items allowed", number)
	}
	
	url := c.BaseUrl + c.NewSuffix
	rep, err := c.GetResource(url)

	err = json.Unmarshal(rep, &top500)
	
	if err != nil {
		return nil, err
	}

	return top500[:number], nil
}

// GetRecentAskStories takes an int number and returns an array of up to number ints that represent the most recent Ask stories
// Constraints: 0 <= number <= 200
func (c *Client) GetRecentAskStories(number int) ([]int, error) {
	var top200 []int
	if number > 200 {
		return top200, fmt.Errorf("Number %d greater than maximum 500 items allowed", number)
	}
	
	url := c.BaseUrl + c.AskSuffix
	rep, err := c.GetResource(url)

	err = json.Unmarshal(rep, &top200)
	
	if err != nil {
		return nil, err
	}

	return top200[:number], nil
}

// GetRecentShowStories takes an int number and returns an array of up to number ints that represent the most recent Show stories
// Constraints: 0 <= number <= 200
func (c *Client) GetRecentShowStories(number int) ([]int, error) {
	var top200 []int
	if number > 200 {
		return top200, fmt.Errorf("Number %d greater than maximum 500 items allowed", number)
	}
	
	url := c.BaseUrl + c.ShowSuffix
	rep, err := c.GetResource(url)

	err = json.Unmarshal(rep, &top200)
	
	if err != nil {
		return nil, err
	}

	return top200[:number], nil
}

// GetRecentJobStories takes an int number and returns an array of up to number ints that represent the most recent Job stories
// Constraints: 0 <= number <= 200
func (c *Client) GetRecentJobStories(number int) ([]int, error) {
	var top200 []int
	if number > 200 {
		return top200, fmt.Errorf("Number %d greater than maximum 500 items allowed", number)
	}
	
	url := c.BaseUrl + c.JobSuffix
	rep, err := c.GetResource(url)

	err = json.Unmarshal(rep, &top200)
	
	if err != nil {
		return nil, err
	}

	return top200[:number], nil
}

func main() {
	c := NewClient()

	story, _ := c.GetItem(8715529)
	fmt.Printf("%+v\n\n", story)

	user, _ := c.GetUser("munrocape")
	fmt.Printf("%+v\n\n", user)

	comment, _ := c.GetItem(8715677)
	fmt.Printf("%+v\n\n", comment)

	top10, _ := c.GetTopStories(10)
	fmt.Printf("%+v\n\n", top10)

	new10, _ := c.GetNewStories(10)
	fmt.Printf("%+v\n\n", new10)

	ask10, _ := c.GetRecentAskStories(10)
	fmt.Printf("%+v\n\n", ask10)

	job10, _ := c.GetRecentJobStories(10)
	fmt.Printf("%+v\n\n", job10)

	show10, _ := c.GetRecentShowStories(10)
	fmt.Printf("%+v\n\n", show10)
}

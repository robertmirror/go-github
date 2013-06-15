// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// EventsService provides access to the event related functions
// in the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/
type EventsService struct {
	client *Client
}

// Event represents a GitHub event.
type Event struct {
	Type      string          `json:"type,omitempty"`
	Public    bool            `json:"public"`
	Payload   json.RawMessage `json:"payload,omitempty"`
	Repo      *Repository     `json:"repo,omitempty"`
	Actor     *User           `json:"actor,omitempty"`
	Org       *Organization   `json:"org,omitempty"`
	CreatedAt *time.Time      `json:"created_at,omitempty"`
	ID        string          `json:"id,omitempty"`
}

// ListPerformedByUser lists the events performed by a user. If publicOnly is
// true, only public events will be returned.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-performed-by-a-user
func (s *EventsService) ListPerformedByUser(user string, publicOnly bool, opt *ListOptions) ([]Event, error) {
	var u string
	if publicOnly {
		u = fmt.Sprintf("users/%v/events/public", user)
	} else {
		u = fmt.Sprintf("users/%v/events", user)
	}

	if opt != nil {
		params := url.Values{
			"page": []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	events := new([]Event)
	_, err = s.client.Do(req, events)
	return *events, err
}

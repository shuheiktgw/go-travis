// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// MessagesService handles communication with the Message
// related methods of the Travis CI API.
type MessagesService struct {
	client *Client
}

// Message represents an individual Message.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/message#message
type Message struct {
	// The message's id
	Id *uint `json:"id"`
	// The message's level
	Level *string `json:"level"`
	// The message's key
	Key *string `json:"key"`
	// The message's code
	Code *string `json:"code"`
	// The message's args
	Args json.RawMessage `json:"args"`
	*Metadata
}

// MessagesOption specifies the optional parameters for messages endpoint
type MessagesOption struct {
	// How many messages to include in the response
	Limit int `url:"limit,omitempty"`
	// How many messages to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
}

// messagesResponse represents a response
// from messages endpoints
type messagesResponse struct {
	Messages []*Message `json:"messages,omitempty"`
}

// ListByRepoId returns a list of messages created by travis-yml for a request, if any exist.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/messages#for_request
func (ms *MessagesService) ListByRepoId(ctx context.Context, repoId uint, requestId uint, opt *MessagesOption) ([]*Message, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/request/%d/messages", repoId, requestId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := ms.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var mr messagesResponse
	resp, err := ms.client.Do(ctx, req, &mr)
	if err != nil {
		return nil, resp, err
	}

	return mr.Messages, resp, err
}

// ListByRepoSlug returns a list of messages created by travis-yml for a request, if any exist.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/messages#for_request
func (ms *MessagesService) ListByRepoSlug(ctx context.Context, repoSlug string, requestId uint, opt *MessagesOption) ([]*Message, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/request/%d/messages", url.QueryEscape(repoSlug), requestId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := ms.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var mr messagesResponse
	resp, err := ms.client.Do(ctx, req, &mr)
	if err != nil {
		return nil, resp, err
	}

	return mr.Messages, resp, err
}

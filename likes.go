package vk_api

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Методы для работы с отметками "Мне нравится"

const (
	TypePost         = "post"
	TypeComment      = "comment"
	TypePhoto        = "photo"
	TypeDocument     = "doc"
	TypeAudio        = "audio"
	TypeVideo        = "video"
	TypeNote         = "note"
	TypePhotoComment = "photo_comment"
	TypeVideoComment = "video_comment"
	TypeTopicComment = "topic_comment"
	TypeSitepage     = "sitepage"
)

type Likes struct {
	Count int         `json:"count"`
	Users []*LikeUser `json:"items"`
}

type LikeUser struct {
	Type      string `json:"profile"`
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// LikesGet получает список идентификаторов пользователей,
// которые добавили заданный объект в свой список «Мне нравится».
//
// Обратите внимание, данные о репостах доступны только для записей,
// созданных текущим пользователем или сообществом, в котором он является администратором.
func (client *VKClient) LikesGet(itemType string, ownerID int, itemID int, count int, params url.Values) (int, []*LikeUser, error) {
	if params == nil {
		params = url.Values{}
	}

	params.Add("type", itemType)
	params.Add("count", strconv.Itoa(count))
	params.Add("owner_id", strconv.Itoa(ownerID))
	params.Add("item_id", strconv.Itoa(itemID))
	params.Add("extended", "1")

	resp, err := client.MakeRequest("likes.getList", params)
	if err != nil {
		return 0, nil, err
	}

	var likes Likes
	json.Unmarshal(resp.Response, &likes)
	return likes.Count, likes.Users, nil
}

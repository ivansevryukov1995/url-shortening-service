package dto

import "github.com/ivansevryukov1995/url-shortening-service/intertnal/model"

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type LinksResponse struct {
	Links []model.Link `json:"links"`
	Count int64        `json:"count"`
}

package main

import "mime/multipart"

type ocrRequest struct {
	File        *multipart.FileHeader `form:"image" json:"-"`
	ImageBase64 string                `json:"image_base64,omitempty"`
	MaskMode    bool                  `json:"mask_mode,omitempty"`
	IsColor     bool                  `json:"is_color,omitempty"`
	SecretMode  bool                  `json:"secret_mode,omitempty"`
}

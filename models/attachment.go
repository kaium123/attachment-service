package models

type AttachmentIDs struct {
	IDs []int `json:"ids"`
}

type Attachment struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type AttachmentResponse struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

type Attachments struct {
	Attachment []Attachment `json:"attachments"`
}

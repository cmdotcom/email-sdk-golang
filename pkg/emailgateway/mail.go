package emailgateway

type ContentType string

const (
	ContentTypePNG  ContentType = "image/png"
	ContentTypeJPG  ContentType = "image/jpg"
	ContentTypeJPEG ContentType = "image/jpeg"
	ContentTypePDF  ContentType = "application/pdf"
	ContentTypeDoc  ContentType = "application/doc"
	ContentTypeDocx ContentType = "application/docx"
	ContentTypeText ContentType = "application/zip"
)

type Address struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Attachment struct {
	Filename string `json:"fileName"`
	// Content is the base64 encoded content of the attachment
	Content     string      `json:"content"`
	ContentType ContentType `json:"contentType"`
	ContentID   string      `json:"contentId"`
}

type Email struct {
	From              Address      `json:"from"`
	ReplyTo           *Address     `json:"replyTo,omitempty"`
	To                []Address    `json:"to"`
	Cc                []Address    `json:"cc,omitempty"`
	Bcc               []Address    `json:"bcc,omitempty"`
	Subject           string       `json:"subject"`
	HTMLBody          string       `json:"html"`
	TextBody          string       `json:"text"`
	Attachments       []Attachment `json:"attachments,omitempty"`
	CustomerReference string       `json:"customerReference,omitempty"`
	Priority          priority     `json:"-"`
}

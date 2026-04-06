package builder

import "github.com/Depth8064/go-cxml/cxml/model"

type Builder struct {
	doc *model.CXML
}

func New() *Builder {
	return &Builder{doc: &model.CXML{}}
}

func (b *Builder) PayloadID(id string) *Builder {
	b.doc.PayloadID = id
	return b
}

func (b *Builder) Timestamp(ts string) *Builder {
	b.doc.Timestamp = ts
	return b
}

func (b *Builder) Version(version string) *Builder {
	b.doc.Version = version
	return b
}

func (b *Builder) From(party *model.Party) *Builder {
	b.doc.From = party
	return b
}

func (b *Builder) To(party *model.Party) *Builder {
	b.doc.To = party
	return b
}

func (b *Builder) Sender(sender *model.Sender) *Builder {
	b.doc.Sender = sender
	return b
}

func (b *Builder) Request(req *model.Request) *Builder {
	b.doc.Request = req
	b.doc.Response = nil
	b.doc.Message = nil
	return b
}

func (b *Builder) Response(resp *model.Response) *Builder {
	b.doc.Response = resp
	b.doc.Request = nil
	b.doc.Message = nil
	return b
}

func (b *Builder) Message(msg *model.Message) *Builder {
	b.doc.Message = msg
	b.doc.Request = nil
	b.doc.Response = nil
	return b
}

func (b *Builder) Status(status *model.Status) *Builder {
	b.doc.Status = status
	return b
}

func (b *Builder) Build() *model.CXML {
	return b.doc
}

func (b *Builder) BuildError(code, text string) *model.CXML {
	b.doc.Response = &model.Response{Status: &model.Status{Code: code, Text: text}}
	b.doc.Request = nil
	b.doc.Message = nil
	return b.doc
}

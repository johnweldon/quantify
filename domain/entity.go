package domain

type Repository interface {
	Save(Entity) (string, error)
	Get(id string) (Entity, error)
	All() ([]Entity, error)
}

type Entity struct {
	ID       string      `json:"entityId",bson:"entityId",xml:"entity-id"`
	Document interface{} `json:"doc",bson:"doc",xml:"entity-document"`
}

func NewEntity(id string, doc interface{}) Entity {
	return Entity{ID: id, Document: doc}
}

func (e Entity) GetID() string            { return e.ID }
func (e Entity) GetDocument() interface{} { return e.Document }

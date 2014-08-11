package quant

type Entity interface {
	GetID() string
}

type Repository interface {
	Save(Entity) error
	Get(id string) (Entity, error)
}

var _ Entity = (*entity)(nil)

type entity struct {
	ID       string      `json:"entityId",bson:"entityId",xml:"entity-id"`
	Document interface{} `json:"doc",bson:"doc",xml:"entity-document"`
}

func NewEntity(id string, doc interface{}) Entity {
	return entity{ID: id, Document: doc}
}

func (e entity) GetID() string { return e.ID }

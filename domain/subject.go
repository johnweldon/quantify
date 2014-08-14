package domain

type SubjectRepository interface {
	Repository
	SaveSubject(s Subject) (string, error)
	GetSubject(name string) (Subject, error)
	AllSubjects() ([]Subject, error)
}

type Subject struct {
	Name string `json:"name",bson:"name",xml:"subject-name"`
}

func NewSubject(name string) Subject {
	return Subject{Name: name}
}

func (s Subject) GetName() string { return s.Name }
func (s Subject) String() string  { return s.Name }

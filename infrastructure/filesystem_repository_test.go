package infrastructure_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/johnweldon/quantify/domain"
	"github.com/johnweldon/quantify/infrastructure"
)

func TestPersistence(t *testing.T) {
	tmp, err := ioutil.TempDir("", "persistence-test")
	if err != nil {
		t.Errorf("could not create temp dir %v", err)
	}
	defer os.RemoveAll(tmp)
	t.Logf("tmpdir: %s", tmp)

	repo, err := infrastructure.NewFilesystemRepository(tmp)
	if err != nil {
		t.Errorf("problem creating repo %v", err)
	}

	entity := domain.NewEntity("", map[string]string{"hello": "world"})

	id, err := repo.Save(entity)
	if err != nil {
		t.Errorf("problem saving entity %v", err)
	}

	if id == "" {
		t.Errorf("repo.Save returned invalid id")
	}

	_, err = repo.Save(domain.NewEntity("", "wow"))
	if err != nil {
		t.Errorf("problem saving a different entity %v", err)
	}

	e, err := repo.Get(id)
	if err != nil {
		t.Errorf("problem restoring entity %v", err)
	}

	m, ok := e.GetDocument().(map[string]interface{})
	if !ok {
		t.Errorf("document of unexpected type %T", e.GetDocument())
	}

	v, ok := m["hello"]
	if !ok {
		t.Errorf("missing expected key hello")
	}

	if v != "world" {
		t.Errorf("missing expected value world")
	}

}

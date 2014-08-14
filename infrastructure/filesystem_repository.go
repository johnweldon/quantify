package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/johnweldon/quant/domain"

	"code.google.com/p/go-uuid/uuid"
)

type filesystemRepository struct {
	Root string
}

var _ domain.Repository = (*filesystemRepository)(nil)

func NewFilesystemRepository(root string) (domain.Repository, error) {
	r, err := filepath.Abs(filepath.Clean(root))
	if err != nil {
		return nil, err
	}

	s, err := os.Stat(r)
	if err != nil {
		return nil, err
	}

	if !s.IsDir() {
		return nil, fmt.Errorf("'%s' is not a directory", root)
	}

	return filesystemRepository{Root: r}, nil
}

func (r filesystemRepository) Save(e domain.Entity) (string, error) {
	id := e.GetID()
	if id == "" {
		id = uuid.New()
		e.ID = id
	}

	target := filepath.Join(r.Root, id)
	if _, err := os.Stat(target); err == nil {
		return "", os.ErrExist
	} else if !os.IsNotExist(err) {
		return "", err
	}

	fd, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	enc := json.NewEncoder(fd)
	err = enc.Encode(&e)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r filesystemRepository) Get(id string) (domain.Entity, error) {
	target := filepath.Join(r.Root, id)
	if _, err := os.Stat(target); err != nil {
		return domain.Entity{}, err
	}

	fd, err := os.Open(target)
	if err != nil {
		return domain.Entity{}, err
	}
	defer fd.Close()

	var entity domain.Entity
	dec := json.NewDecoder(fd)
	err = dec.Decode(&entity)
	if err != nil {
		return domain.Entity{}, err
	}

	return entity, nil
}

func (r filesystemRepository) All() ([]domain.Entity, error) {
	entities := []domain.Entity{}
	var fn filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		if path != r.Root && info.IsDir() {
			return filepath.SkipDir
		}
		if entity, err := r.Get(info.Name()); err == nil {
			entities = append(entities, entity)
		}
		return nil
	}
	filepath.Walk(r.Root, fn)
	return entities, nil
}

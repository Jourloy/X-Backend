package node_rep

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[node-database]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.NodeRepository

type NodeRepository struct {
	db gorm.DB
}

// Init создает репозиторий узла
func Init() {
	go migration()

	Repository = &NodeRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := storage.Database.AutoMigrate(
		&repositories.Node{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}

// Create создает узел
func (r *NodeRepository) Create(create *repositories.NodeCreate) (*repositories.Node, error) {
	node := repositories.Node{
		ID:        uuid.NewString(),
		X:         create.X,
		Y:         create.Y,
		Difficult: create.Difficult,
		Walkable:  create.Walkable,
		SectorID:  create.SectorID,
	}

	res := r.db.Create(&node)
	if res.Error != nil {
		return nil, res.Error
	}

	return &node, nil
}

// GetOne возвращает первый узел, попавший под условие
func (r *NodeRepository) GetOne(query *repositories.NodeGet) (*repositories.Node, error) {
	node := repositories.Node{}

	res := r.db.First(&node, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &node, nil
}

// GetAll возвращает все рынки
func (r *NodeRepository) GetAll(query *repositories.NodeGet) (*[]repositories.Node, error) {
	nodes := []repositories.Node{}

	res := r.db.Find(&nodes, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &nodes, nil
}

// UpdateOne обновляет рынок
func (r *NodeRepository) UpdateOne(node *repositories.Node) error {
	res := r.db.Save(&node)
	return res.Error
}

// DeleteOne удаляет рынок
func (r *NodeRepository) DeleteOne(node *repositories.Node) error {
	res := r.db.Delete(&node, node)
	return res.Error
}

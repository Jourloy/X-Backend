package node_rep

import (
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
func (r *NodeRepository) Create(node *repositories.Node) {
	node.ID = uuid.NewString()
	r.db.Create(&node)
}

// GetOne возвращает первый узел, попавший под условие
func (r *NodeRepository) GetOne(node *repositories.Node) {
	r.db.First(&node, node)
}

// GetAll возвращает все рынки
func (r *NodeRepository) GetAll(dest *[]repositories.Node, query repositories.NodeGetAll) {
	node := repositories.Node{
		X:         *query.X,
		Y:         *query.Y,
		Walkable:  *query.Walkable,
		Difficult: *query.Difficult,
		SectorID:  query.SectorID,
	}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(node).Limit(limit).Find(&dest)
}

// UpdateOne обновляет рынок
func (r *NodeRepository) UpdateOne(node *repositories.Node) {
	r.db.Save(&node)
}

// DeleteOne удаляет рынок
func (r *NodeRepository) DeleteOne(node *repositories.Node) {
	r.db.Delete(&node)
}

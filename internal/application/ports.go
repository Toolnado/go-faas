package application

import (
	"context"

	"github.com/Toolnado/go-faas/internal/domain"
	"github.com/google/uuid"
)

type IFunctionRepository interface {
	Insert(ctx context.Context, f *domain.Function) (*domain.Function, error)
	Get(ctx context.Context, id uuid.UUID) (*domain.Function, error)
}

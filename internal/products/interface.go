package products

import (
	"context"

	"github.com/EugeniaB8315/Go_C14_S_DB/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, producto domain.Producto) (domain.Producto, error)
	// recibe un ctx y devuelve un slice(array/lista)product junto al error
	GetAll(ctx context.Context) ([]domain.Producto, error)
	GetByID(ctx context.Context, id int) (domain.Producto, error)
	Update(ctx context.Context, producto domain.Producto, id int) (domain.Producto, error)
	Delete(ctx context.Context, id int) error
}

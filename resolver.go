package products

import (
	"context"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/zerbitx/federation-demo/products/ptr"
)

type (
	ProductService interface {
		ByUPC(string) *Product
		TopProducts(ctx context.Context, first *int) ([]*Product, error)
	}
	
	Resolver struct {
		ProductService
	}
	
	queryResolver struct {
		*Resolver
	}
)

const (
	typeNameKey     = "__typename"
	productTypeName = "Product"
	upcKey          = "upc"
)

func New(service ProductService) *Resolver {
	return &Resolver{
		ProductService: service,
	}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) _entities(ctx context.Context, representations []map[string]interface{}) ([]_Entity, error) {
	var entities []_Entity

	for _, rep := range representations {
		typeName, ok := rep[typeNameKey]

		if !ok {
			return nil, errors.New("__typename required")
		}

		if typeName != productTypeName {
			return nil, errors.New("Invalid type name, only Products is supported")
		}

		if upc := r.ByUPC(rep[upcKey].(string)); upc != nil {
			entities = append(entities, upc)
		}
	}

	return entities, nil
}

func (r *queryResolver) _service(ctx context.Context) (*_Service, error) {
	schema, err := ioutil.ReadFile("./base_schema.graphql")

	if err != nil {
		return nil, err
	}

	return &_Service{
		Sdl: ptr.String(string(schema)),
	}, nil
}

func (r *queryResolver) TopProducts(ctx context.Context, first *int) ([]*Product, error) {
	return r.ProductService.TopProducts(ctx, first)
}

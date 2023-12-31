package schema

import (
	"encoding/json"
	"fmt"

	// "github.com/tiagoposse/connect/ent/hook"
	// "github.com/tiagoposse/connect/ent/user"

	"github.com/tiagoposse/connect/internal/types"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen"

	authz "github.com/tiagoposse/go-auth/authorization"
	ogauthz "github.com/tiagoposse/ogent-auth/authorization"
)

// ApiKey holds the schema definition for the ApiKey entity.
type ApiKey struct {
	ent.Schema
}

// Annotations of the ApiKey.
func (ApiKey) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		ogauthz.WithAllScopes(types.AdminAll, types.UserAll, types.UserApiKeyWrite),
	}
}

// Fields of the ApiKey.
func (ApiKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Immutable(),
		field.String("key").Immutable().Sensitive(),
		field.Other("scopes", authz.Scopes{}).
			SchemaType(map[string]string{dialect.Postgres: "varchar"}).
			Default(authz.Scopes{types.UserAll}).
			Immutable().
			Annotations(
				entoas.Schema(ogen.String().
					AsEnum(nil,
						json.RawMessage(fmt.Sprintf(`"%s"`, types.UserAll)),
						json.RawMessage(fmt.Sprintf(`"%s"`, types.AdminAll)),
					).
					AsArray(),
				),
			),
		field.String("user_id").Immutable(),
	}
}

// Edges of the ApiKey.
func (ApiKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Unique().Ref("keys").Immutable().Required().Field("user_id"),
	}
}

// Hooks of the ApiKey.
// func (ApiKey) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		hook.On(
// 			func(next ent.Mutator) ent.Mutator {
// 				return hook.ApiKeyFunc(func(ctx context.Context, m *gen.ApiKeyMutation) (ent.Value, error) {
// 					name, _ := m.Name()

// 					keys, err := m.Client().ApiKey.Query().WithUser(func(q *gen.UserQuery) {
// 						uID, _ := m.UserID()
// 						q.Where(user.IDEQ(uID))
// 					}).All(ctx)
// 					if err != nil {
// 						return nil, err
// 					}

// 					if len(keys) >= 2 {
// 						return nil, utils.NewMaxItemsError("api keys for user")
// 					}

// 					for _, k := range keys {
// 						if k.Name == name {
// 							return nil, utils.NewMaxItemsError("api key with same name for user")
// 						}
// 					}

// 					return next.Mutate(ctx, m)
// 				})
// 			},
// 			// Limit the hook only for these operations.
// 			ent.OpCreate,
// 		),
// 	}
// }

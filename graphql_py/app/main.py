from ariadne import make_executable_schema, load_schema_from_path
from ariadne.asgi import GraphQL
from app.schema import type_defs
from  app.resolvers import resolver
from app.database import engine, Base

Base.metadata.create_all(bind=engine)

# query = {
#     "users": queries.resolve_users,
#     "user": queries.resolve_user,
#     "posts": queries.resolve_users,
#     "post": queries.resolve_post,
# }
#
# mutations = {
#     "createUser": mutation.resolve_create_user,
#     "createPost": mutation.resolve_create_post,
# }

schema = make_executable_schema(type_defs, *resolver)

app = GraphQL(schema=schema, debug=True)
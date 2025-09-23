from ariadne import convert_kwargs_to_snake_case, QueryType
from app.models import Post, User
from app.database import get_db

query = QueryType()

@query.field("users")
@convert_kwargs_to_snake_case
def resolve_users(obj, info):
    db = next(get_db())
    users = db.query(User).all()
    return users

@query.field("user")
@convert_kwargs_to_snake_case
def resolve_user(obj, info, id):
    db = next(get_db())
    user = db.query(User).filter(User.id == id).first()
    return user

@query.field("post")
@convert_kwargs_to_snake_case
def resolve_post(obj, info, id):
    db = next(get_db())
    post = db.query(Post).filter(Post.id == id).first()
    return post

@query.field("posts")
@convert_kwargs_to_snake_case
def resolve_posts(obj, info):
    db = next(get_db())
    posts = db.query(Post).all()
    return posts
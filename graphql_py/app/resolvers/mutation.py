from ariadne import convert_kwargs_to_snake_case, MutationType
from app.models import User, Post
from app.database import get_db

mutation = MutationType()

@mutation.field("createPost")
@convert_kwargs_to_snake_case
def resolve_create_post(obj, info, input):
    db = next(get_db())
    try:
        post = Post(
            title=input["title"],
            content=input["content"],
            author_id=input["author_id"],
        )
        db.add(post)
        db.commit()
        db.refresh(post)
        return post
    except Exception as e:
        db.rollback()
        return {
            "detail": False,
            "error": str(e)
        }


@mutation.field("createUser")
@convert_kwargs_to_snake_case
def resolve_create_user(obj, info, input):
    db = next(get_db())
    try:
        user = User(
            name=input["name"],
            email=input["email"],
            age=input["age"],
        )
        db.add(user)
        db.commit()
        db.refresh(user)
        return user
    except Exception as e:
        db.rollback()
        return {
            "detail": True,
            "error": str(e)
        }

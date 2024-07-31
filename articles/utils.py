import random
from django.utils.text import slugify

def slugify_article_instance_title(instance, save=False, new_slug=None):
    """
    Create a unique slug for each new article created.
    """
    if new_slug is not None:
        slug = new_slug
    else:
        slug = slugify(instance.title)
    Klass = instance.__class__
    qs = Klass.objects.filter(slug=slug).exclude(id=instance.id)
    # if that slug already exists in the database
    if qs.exists():
        rand_int = random.randint(300_000, 500_000)
        new_slug = f"{slug} - {rand_int}"
        slug = new_slug
        return slugify_article_instance_title(instance, save=save, new_slug=slug)
    instance.slug = slug
    if save:
        instance.save()
    return instance

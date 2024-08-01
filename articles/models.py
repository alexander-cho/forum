from django.db import models
from django.db.models.signals import pre_save, post_save
from django.utils import timezone

from .utils import slugify_article_instance_title

# Create your models here.


class Article(models.Model):
    title = models.CharField(max_length=100)
    slug = models.SlugField(unique=True, null=True, blank=True)
    content = models.TextField()
    timestamp = models.DateTimeField(auto_now_add=True)
    last_updated = models.DateTimeField(auto_now=True)
    publish = models.DateField(auto_now_add=False, auto_now=False, default=timezone.now, null=True, blank=True)

    def save(self, *args, **kwargs):
        # if self.slug is None:
        #     slugify_article_instance_title(self, save=False)
        super().save(*args, **kwargs)


def article_pre_save(sender, instance, *args, **kwargs):
    if instance.slug is None:
        slugify_article_instance_title(instance, save=False)

pre_save.connect(article_pre_save, sender=Article)


def article_post_save(sender, instance, created, *args, **kwargs):
    if created:
        slugify_article_instance_title(instance, save=True)

post_save.connect(article_post_save, sender=Article)

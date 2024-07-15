from django.contrib import admin

# Register your models here.
from .models import Article


# what to display in admin interface
class ArticleAdmin(admin.ModelAdmin):
    list_display = ['title']


admin.site.register(Article, ArticleAdmin)

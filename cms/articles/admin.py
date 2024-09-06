from django.contrib import admin

# Register your models here.
from .models import Article


# what to display in admin interface
class ArticleAdmin(admin.ModelAdmin):
    list_display = ['id', 'title', 'slug', 'timestamp', 'last_updated']
    search_fields = ['title', 'content']


admin.site.register(Article, ArticleAdmin)

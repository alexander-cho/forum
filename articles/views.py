from django.shortcuts import render

from .models import Article

# Create your views here.


def article_detail_view(request, id=None):
    article_object = None
    if id is not None:
        article_object = Article.objects.get(id=id)
    context = {
        'object': article_object
    }
    return render(request, 'articles/detail.html', context=context)

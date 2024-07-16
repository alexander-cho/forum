from django.shortcuts import render, redirect

from .models import Article

# Create your views here.


def article_search_view(request):
    query_dict = request.GET
    # <input type="text" name="q"/>
    try:
        query = int(query_dict.get('q'))
    except ValueError:
        query = None
    article_object = None
    if query is not None:
        article_object = Article.objects.get(id=query)
    context = {
        'article_object': article_object
    }
    return render(request, 'articles/search.html', context=context)


def article_detail_view(request, id=None):
    article_object = None
    if id is not None:
        article_object = Article.objects.get(id=id)
    context = {
        'article_object': article_object
    }
    return render(request, 'articles/detail.html', context=context)

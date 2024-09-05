from django.shortcuts import render, redirect
from django.contrib.auth.decorators import login_required
from django.http import Http404

from .models import Article
from .forms import ArticleCreationForm

# Create your views here.


def article_search_view(request):
    print(request.GET)
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


def article_detail_view(request, slug=None):
    article_object = None
    if slug is not None:
        try:
            article_object = Article.objects.get(slug=slug)
        except Article.DoesNotExist:
            raise Http404
        except Article.MultipleObjectsReturned:
            article_object = Article.objects.get(slug=slug).first()
        except:
            raise Http404
    context = {
        'article_object': article_object
    }
    return render(request, 'articles/detail.html', context=context)


@login_required
def article_create_view(request):
    form = ArticleCreationForm(request.POST or None)
    context = {
        'form': form
    }
    if form.is_valid():
        # validate form using ModelForm class
        article_object = form.save()
        context['article_object'] = article_object
        context['created'] = True
        return redirect(article_object.get_absolute_url())

    return render(request, 'articles/create.html', context=context)

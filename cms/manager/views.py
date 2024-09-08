from django.http import HttpResponse
from django.template.loader import render_to_string

from articles.models import Article


def home_view(request):
    articles = Article.objects.all()
    context = {
        'articles': articles
    }
    html_string = render_to_string('home-view.html', context)
    return HttpResponse(html_string)

"""
Return html
"""

from django.http import HttpResponse
from django.template.loader import render_to_string

from articles.models import Article


def home_view(request):
    # article_object = Article.objects.get(id=1)
    articles = Article.objects.all()
    context = {
        'articles': articles,
        # 'id': article_object.id,
        # 'title': article_object.title,
        # 'content': article_object.content,
    }
    html_string = render_to_string('home-view.html', context=context)
    # html_string = """
    # <h1>{title} (id: {id})</h1>
    # <p>{content}</p>
    # """.format(**context)
    return HttpResponse(html_string)

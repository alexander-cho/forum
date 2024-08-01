from django.test import TestCase

from .models import Article
from .utils import slugify

# Create your tests here.
class TestArticle(TestCase):
    """
    Tests for the Article model.
    """
    def setUp(self):
        self.number_of_articles = 10
        for _ in range(0, self.number_of_articles):
            Article.objects.create(title='hello world', content='asdfcsea')

    def test_queryset_exists(self):
        qs = Article.objects.all()
        self.assertTrue(qs.exists())

    def test_queryset_count(self):
        qs = Article.objects.all()
        self.assertEqual(qs.count(), self.number_of_articles)

    def test_queryset_valid_slug(self):
        obj = Article.objects.all().order_by('id').first()
        title = obj.title
        slug_title = slugify(title)
        slug = obj.slug
        self.assertEqual(slug, slug_title)

    def test_queryset_unique_slug(self):
        """
        test for unique slugs, which will contain the random number 'tail'
        """
        qs = Article.objects.exclude(slug__iexact='hello-world')
        for obj in qs:
            title = obj.title
            slug_title = slugify(title)
            slug = obj.slug
            self.assertNotEqual(slug, slug_title)

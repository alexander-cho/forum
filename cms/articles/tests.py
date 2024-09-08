from django.test import TestCase
from django.utils.text import slugify
from .models import Article
from .utils import slugify_article_instance_title


class TestArticle(TestCase):
    """
    Tests for the Article model.
    """
    def setUp(self):
        """
        set up for each test for consistent state, create Article objects
        """
        self.number_of_articles = 100
        for _ in range(0, self.number_of_articles):
            Article.objects.create(title='hello world', content='asdfcsea')

    def test_queryset_exists(self):
        """
        test that queryset for Article exists; returns something
        """
        qs = Article.objects.all()
        self.assertTrue(qs.exists())

    def test_queryset_count(self):
        """
        test that correct number of Article objects has been created
        """
        qs = Article.objects.all()
        self.assertEqual(qs.count(), self.number_of_articles)

    def test_queryset_valid_slug(self):
        """
        test slug of Article object matches slugified version of its title
        """
        obj = Article.objects.all().order_by('id').first()
        title = obj.title
        slug_title = slugify(title)
        slug = obj.slug
        self.assertEqual(slug, slug_title)

    def test_queryset_unique_slug(self):
        """
        test for unique slugs, which will contain the random number ending
        """
        qs = Article.objects.exclude(slug__iexact='hello-world')
        for obj in qs:
            title = obj.title
            slug_title = slugify(title)
            slug = obj.slug
            self.assertNotEqual(slug, slug_title)

    def test_slugify_article_instance_title(self):
        """
        test utility function for generating unique slugs
        """
        obj = Article.objects.all().last()
        new_slugs = []
        for _ in range(0, 25):
            instance = slugify_article_instance_title(obj, save=False)
            new_slugs.append(instance.slug)
        unique_slugs = list(set(new_slugs))
        self.assertEqual(len(new_slugs), len(unique_slugs))

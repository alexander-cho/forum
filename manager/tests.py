import os

from django.test import TestCase
from django.contrib.auth.password_validation import validate_password


class ManagerConfigTest(TestCase):
    def test_secret_key_strength(self):
        SECRET_KEY = os.environ.get('DJANGO_SECRET_KEY_TEST')
        try:
            is_strong = validate_password(SECRET_KEY)
        except Exception as e:
            message = f'Bad secret key {e.messages}'
            self.fail(message)

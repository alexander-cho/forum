from django import forms
from .models import Article


class ArticleCreationForm(forms.ModelForm):
    class Meta:
        model = Article
        fields = ['title', 'content']

    def clean(self):
        data = self.cleaned_data
        title = data.get('title')
        queryset = Article.objects.all().filter(title__icontains=title)
        if queryset.exists():
            self.add_error('title', f'that title: {title} already exists, choose another one.')
        return data


class ArticleCreationFormOld(forms.Form):
    title = forms.CharField(max_length=100)
    content = forms.CharField(max_length=1000)
    
    def clean(self):
        cleaned_data = self.cleaned_data
        print('all clean', cleaned_data)
        title = cleaned_data.get('title')
        if title.lower().strip() == 'all clean':
            raise forms.ValidationError('that title is already taken.')
        return cleaned_data
    
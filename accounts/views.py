from django.contrib.auth import authenticate, login
from django.shortcuts import render, redirect

# Create your views here.


def login_view(request):
    context = {}
    if request.method == 'POST':
        username = request.POST.get('username')
        password = request.POST.get('password')
        print(username, password)
        user = authenticate(request, username=username, password=password)
        if user is None:
            context = {'error': 'error!'}
            return render(request, 'accounts/login.html', context=context)
        login(request, user)
        return redirect('/')
    return render(request, 'accounts/login.html', context=context)


def logout_view(request):
    context = {}
    return render(request, 'accounts/logout.html', context=context)


def register_view(request):
    context = {}
    return render(request, 'accounts/register.html', context=context)
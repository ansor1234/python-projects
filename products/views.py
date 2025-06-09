from django.http import HttpResponse
# from django.views import View
from django.shortcuts import render


def index(request):
    return HttpResponse("Hello, world. You're at the products index.")


def new(request2):
    return HttpResponse("Hello, Dunyo. You're at the products/new index.")
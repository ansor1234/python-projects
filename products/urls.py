from django.urls import path
from . import views
from . import views2





urlpatterns = [
    path('', views.index),
    path('', views2.index2)
]
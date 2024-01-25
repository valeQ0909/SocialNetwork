from django.urls import path
from . import views

urlpatterns = [
    path('mnist', views.DigitalMnist)
]
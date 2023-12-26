from django.urls import path

from . import views

urlpatterns =[
    path('<str:category>', views.PostViews.as_view())
]
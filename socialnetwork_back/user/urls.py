from django.urls import path
from . import views

urlpatterns = [
    path('<str:account>', views.UserView.as_view()),
    path('<str:account>/avatar', views.users_views),
    path('info/', views.getUserInfo)
]
"""socialnetwork_back URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/3.2/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path,include
from django.conf.urls.static import static
from django.conf import settings
from user import views as user_views

urlpatterns = [
    path('admin/', admin.site.urls),
    path('user', user_views.UserView.as_view()),  # 这里需要将UserView类转换为as_view()方法
    path('user/', include('user.urls')),
    path('tokens/', include('usertoken.urls')),
    path('topics/', include('topic.urls')),
    path('post/', include('post.urls')),
    path('messages/', include('message.urls')),
]

urlpatterns += static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)

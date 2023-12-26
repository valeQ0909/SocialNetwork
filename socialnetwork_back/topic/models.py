from django.db import models
from user.models import User

class Topic(models.Model):
    category = models.CharField(verbose_name="文章分类", max_length=20)
    limit = models.CharField(verbose_name="文章权限", max_length=10)
    introduce = models.CharField(verbose_name="文章简介", max_length=90)
    content = models.TextField(verbose_name="文章内容")
    created_time = models.DateTimeField(auto_now_add=True)
    updated_time = models.DateTimeField(auto_now=True)
    author = models.ForeignKey(User, on_delete=models.CASCADE)

from django.db import models
from user.models import User
from topic.models import Topic
# Create your models here.

class Message(models.Model):
    content = models.CharField(max_length=100)
    created_time = models.DateTimeField(auto_now_add=True)
    parent_message = models.IntegerField(verbose_name='回复的留言id')
    publisher = models.ForeignKey(User, on_delete=models.CASCADE)
    topic = models.ForeignKey(Topic, on_delete=models.CASCADE)


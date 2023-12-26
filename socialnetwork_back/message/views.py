import json
from django.http import JsonResponse
from message.models import Message
from tools.logging_dec import logging_check
from topic.models import Topic

#异常码10400-10499

# Create your views here.
@logging_check
def message_view(request, topic_id):
    user = request.myuser
    json_str = request.body
    json_obj = json.loads(json_str)
    content = json_obj['content']
    parent_id = json_obj.get('parent_id', 0)  #消息的父消息， 没有这个消息则给一个0

    try:
        topic = Topic.objects.get(id=topic_id)
    except Exception as e:    #用户有删除自己文章的功能，所以要检验一下
        result = {'code':10400, 'error':'Thee topic is not existed'}
        return JsonResponse(result)

    Message.objects.create(topic=topic, content=content, parent_message=parent_id, publisher=user)
    return JsonResponse({'code':200})

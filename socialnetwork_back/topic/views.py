from django.views import View
from django.http import JsonResponse
from django.utils.decorators import method_decorator
from tools.logging_dec import logging_check, get_user_by_request
from tools.cache_dec import cache_set
import json
from .models import Topic
from user.models import User
from django.core.cache import cache
from message.models import Message

#异常码 10300-10399

# Create your views here.
class TopicViews(View):

    #删除缓存
    def clear_topics_caches(self, request):
        #先取出路由
        path = request.path_info
        #前缀
        cache_key_p = ['topics_cache_self_','topics_cache_']
        #后缀
        cache_key_h = ['', '?category=龙子湖畔', '?category=二手市场']
        all_keys = []
        for key_p in cache_key_h:
            for key_h in cache_key_h:
                all_keys.append(key_p + key_h)

        cache.delete_many(all_keys)

    #帖子详情页
    def make_topic_res(self, author, author_topic):

        # 关联留言和回复
        all_messages = Message.objects.filter(topic=author_topic).order_by('-created_time')

        msg_list = []
        m_count = 0
        rep_dic = {}
        for msg in all_messages:
            if msg.parent_message:
                # 回复
                rep_dic.setdefault(msg.parent_message, [])
                rep_dic[msg.parent_message].append({'msg_id': msg.id, 'publisher': msg.publisher.username,
                                                    'publisher_avatar': "http://127.0.0.1:8000/media/" + str(msg.publisher.avatar),
                                                    'content': msg.content,
                                                    'created_time': msg.created_time.strftime("%Y-%m-%d %H:%M:%S")})
            else:
                # 留言
                m_count += 1
                msg_list.append({'id': msg.id, 'content': msg.content, 'publisher': msg.publisher.username,
                                 'publisher_avatar':"http://127.0.0.1:8000/media/" + str(msg.publisher.avatar),
                                 'created_time': msg.created_time.strftime("%Y-%m-%d %H:%M:%S"),
                                 'reply': []})

        for m in msg_list:
            if m['id'] in rep_dic:
                m['reply'] = rep_dic[m['id']]

        res = {'code':200, 'data':{}}
        res['data']['avatar'] = "http://127.0.0.1:8000/media/" + str(author.avatar)
        res['data']['username'] = author.username
        res['data']['category'] = author_topic.category
        res['data']['created_time'] = author_topic.created_time.strftime("%Y-%m-%d %H:%M:%S")
        res['data']['content'] = author_topic.content
        res['data']['introduce'] = author_topic.introduce
        res['data']['author'] = author.username
        res['data']['messages'] = msg_list
        res['data']['messages_count'] = m_count
        return res

    #帖子列表页
    def make_topics_res(self, author, author_topics):

        res = {'code':200, 'data':{}}
        topics_res = []
        for topic in author_topics:
            d = {}
            d['account'] = author.id
            d['id'] = topic.id
            d['category'] = topic.category
            d['created_time'] = topic.created_time.strftime("%Y-%m-%d %H:%M:%S")
            d['introduce'] = topic.introduce
            d['author'] = author.username
            d['avatar'] = "http://127.0.0.1:8000/media/" + str(author.avatar)
            topics_res.append(d)

        res['data']['topics'] = topics_res
        return res

    @method_decorator(logging_check)
    def post(self, request, author_id):

        author = request.myuser
        #提取出前端数据
        json_str = request.body
        json_obj = json.loads(json_str)
        content = json_obj['content']
        content_text = json_obj['content_text']
        introduce = content_text[:80]
        limit = json_obj['limit']
        category = json_obj['category']

        if limit not in ['public','anonymous']:
            result = {'code':10300, 'error':'The limit error ~'}
            return JsonResponse(result)

        #创建topic数据
        Topic.objects.create(content=content, limit=limit, category=category, introduce=introduce,
                             author=author)

        #删除缓存
        self.clear_topics_caches(request)

        return JsonResponse({'code':200})

    @method_decorator(cache_set(30))  # 存30s
    def get(self, request, author_id):
        print('----view---in-----')
        #访问者visitor
        #当前被访问用户author
        try:
            author = User.objects.get(account=author_id)
        except Exception as e:
            result = {'code':10301, 'error':'The account is not existed'}
            return JsonResponse(result)

        visitor = get_user_by_request(request)
        visitor_account = None
        if visitor:
            visitor_account = visitor.id

        #/topics/001?t_id=1
        t_id = request.GET.get('t_id')
        if t_id:
            #获取指定文章数据
            t_id = int(t_id)
            if visitor_account == author_id:
                try:
                    author_topic = Topic.objects.get(id=t_id, author_id=author_id)
                except Exception as e:
                    result = {'code':10302, 'error':'No Topic'}
                    return JsonResponse(result)
            else:
                try:
                    author_topic = Topic.objects.get(id=t_id, author_id=author_id, limit='public')
                except Exception as e:
                    result = {'code':10303, 'error':'No Topic'}
                    return JsonResponse(result)

            res = self.make_topic_res(author, author_topic)
            return JsonResponse(res)

        else:
            #获取列表页信息
            category = request.GET.get('category')

            # /topics/001?category=[]
            if category in ['龙子湖畔', '二手市场', '失物招领', '选课指南', '表白墙', '考研', '招聘']:
                if visitor_account == author_id:
                    # 博主访问自己的博客
                    author_topics = Topic.objects.filter(author_id=author_id, category=category)
                else:
                    # 访客在访问
                    author_topics = Topic.objects.filter(author_id=author_id, limit='public', category=category)

            # /topics/001
            else:
                if visitor_account == author_id:
                    # 博主访问自己的博客
                    author_topics = Topic.objects.filter(author_id=author_id)
                else:
                    # 访客在访问
                    author_topics = Topic.objects.filter(author_id=author_id, limit='public')

                res = self.make_topics_res(author, author_topics)

            return JsonResponse(res)

from django.views import View
from django.http import JsonResponse
from topic.models import Topic
# Create your views here.

#异常码10500-10599
class PostViews(View):

    def make_posts_res(self, posts):
        res = {'code': 200, 'data': {}}
        posts_res = []
        for post in posts:
            d = {}
            d['account'] = post.author.account
            d['id'] = post.id
            d['category'] = post.category
            d['created_time'] = post.created_time.strftime("%Y-%m-%d %H:%M:%S")
            d['introduce'] = post.introduce
            d['author'] = post.author.username
            d['avatar'] = "http://127.0.0.1:8000/media/" + str(post.author.avatar)
            posts_res.append(d)

        res['data']['posts'] = posts_res
        return res

    def get(self, request, category):

        category = request.GET.get('category')
        if category:
            try:
                posts = Topic.objects.filter(category=category, limit='public')
            except Exception as e:
                result = {'code': 10500, 'error': 'the category is not existed'}
                return JsonResponse(result)
        else:
            posts = Topic.objects.all()

        res = self.make_posts_res(posts)
        return JsonResponse(res)
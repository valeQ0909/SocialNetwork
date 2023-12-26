from django.views import View
from django.http import JsonResponse
import json
from .models import User
import hashlib
import os
from django.utils.decorators import method_decorator
from tools.logging_dec import logging_check

#django提供了一个装饰器 method_decorator


def get_type(str):
    ans = ''
    for i in reversed(str):
        if i == '.':
            ans += i
            break
        else:
            ans += i
    ans = ans[::-1]
    return ans

@logging_check
def users_views(request, account):
    if request.method != 'POST':
        result = {'code':10103, 'error':'Please use POST'}
        return JsonResponse(result)

    user = request.myuser

    avatar = request.FILES['avatar']
    avatar.name = account + get_type(avatar.name)
    base_dir = "./media/images/avatar/"
    final_dir = base_dir + avatar.name
    if os.path.isfile(final_dir):
        os.remove(final_dir)
    user.avatar = avatar
    user.save()
    return JsonResponse({'code':200, 'url':'images/avatar/'+avatar.name})


# CBV
# 更灵活【可继承】
# 对未定义的http method请求 直接返回405响应
class UserView(View):
    def get(self, request, account=None):
        if account:
            try:
                user = User.objects.get(account=account)
            except Exception as e:
                result = {'code':10102, 'error':'The account is wrong ~~~'}
                return JsonResponse(result)
            result = {'code':200, 'account':account, 'data':{
    'info':user.info, 'sign':user.sign, 'username':user.username, 'avatar':str(user.avatar)
            }}
            return JsonResponse(result)
        else:
            pass
        return

    def post(self, request):
        json_str = request.body  # 前端传来的是json字符串
        json_obj = json.loads(json_str) #通过json.load将字符串转为python字典
        account = json_obj['account']
        password_1 = json_obj['password_1']
        password_2 = json_obj['password_2']

        # 参数基本检查
        if password_1 != password_2:
            # 异常码 10100 - 10199 由具体公司的老大来决定
            result = {'coded':10100, 'error':'The password is not same ~~~'}
            return JsonResponse(result)

        # 检查账号是否可用
        old_user = User.objects.filter(account=account)
        # filter 会返回一组kv set，但get会返回唯一一个值
        if old_user:
            result = {'coded':10101, 'error':'The username is already existed'}
            return result

        # 向User插入数据（密码用md5存储）
        p_m = hashlib.md5()
        p_m.update(password_1.encode()) #因为password是一个字符串，所以要先encode转换一下

        User.objects.create(account=account, username="用户"+account, password=p_m.hexdigest(),avatar="images/avatar/default.jpg")

        result = {'code': 200, 'account': account, 'data':{"avatar":"images/avatar/default.jpg"}}
        return JsonResponse(result)

    @method_decorator(logging_check)
    def put(self, request, account=None):
        #更新用户数据
        json_str = request.body
        json_obj = json.loads(json_str)

        user = request.myuser

        user.username = json_obj['username']
        user.sign = json_obj['sign']
        user.info = json_obj['info']

        user.save()
        return JsonResponse({'code':200, 'data':{'username':user.username}})
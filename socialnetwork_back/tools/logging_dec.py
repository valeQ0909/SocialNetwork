from django.http import JsonResponse
from django.conf import settings
import jwt
from user.models import User

def logging_check(func):
    def wrap(request, *args, **kwargs):
        #获取token request.META.get('HTTP_AUTHORIZATION')  注：django为了处理和wsgi重名的问题，将请求头里面的属性都重命名了
        token = request.META.get('HTTP_AUTHORIZATION')
        if not token:
            result = {'code':403, 'error':'Please login'}
            return JsonResponse(result)

        #校验jwt
        try:
            res = jwt.decode(token, settings.JWT_TOKEN_KEY, algorithms='HS256')
        except Exception as e:
            print('jwt decode error is %s'%(e))
            result = {'code':403, 'error':'Please login'}
            return JsonResponse(result)

        #失败， coded 403 error:please login

        account = res['username']
        user = User.objects.get(account=account)
        request.myuser = user

        return func(request, *args, **kwargs)
    return wrap


def get_user_by_request(request):
    #尝试性获取登录用户
    # return User obj or NONE
    token = request.META.get('HTTP_AUTHORIZATION')
    if not token:
        return None
    try:
        res = jwt.decode(token, settings.JWT_TOKEN_KEY, algorithms='HS256')
    except Exception as e:
        return None

    account = res['username']
    user = User.objects.get(account=account)
    return user

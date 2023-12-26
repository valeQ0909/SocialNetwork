import time
import jwt
from django.http import JsonResponse
import json
from user.models import User
import hashlib
from django.conf import settings

# 异常码10200-10299
# Create your views here.
def tokens_view(request):
    if request.method != 'POST':
        result = {'code':10200, 'error':'Please use POST !'}
        return JsonResponse(result)
    json_str = request.body
    json_obj = json.loads(json_str)
    account = json_obj['account']
    password = json_obj['password']

    #校验用户名和密码
    try:
        user = User.objects.get(account=account)
    except Exception as e:
        result = {'code':10201,'error':'The username is wrong'}
        return JsonResponse(result)

    p_m = hashlib.md5()
    p_m.update(password.encode())
    if p_m.hexdigest() != user.password:
        result = {'code':10201, 'error':'The username or password is wrong'}
        return JsonResponse(result)

    #记录会话状态
    token = make_token(account)
    result = {'code':200, 'account':account, 'data':{'token':token, 'avatar':str(user.avatar),  'username':user.username}} #这里的token是个字节串，所以需要decode一下
    return JsonResponse(result)                                      #但是貌似2.x版本以后自动就是字符串了，不用decoded了

def make_token(account, expire=3600*24):
    key = settings.JWT_TOKEN_KEY
    now_t = time.time()
    payload_data = {'username': account, 'exp': now_t+expire}
    return jwt.encode(payload_data, key, algorithm='HS256')

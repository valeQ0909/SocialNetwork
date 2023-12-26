from django.db import models
import random

# Create your models here.

MEDIR_ADDR = "http://localhost:8000/media/"

def default_sign():
    signs = ['我见青山多妩媚，料青山见我应如是', '且放白鹿青崖间，须行即骑访名山', '仿佛兮若轻云之蔽月，飘摇兮若流风之回雪']
    return random.choices(signs)[0]

class User(models.Model):
    account = models.CharField(verbose_name='用户账号-用户唯一标识', max_length=50, primary_key=True)  # 主键
    username = models.CharField(verbose_name="用户名", max_length=30,default="用户"+str(id))
    password = models.CharField(verbose_name="密码", max_length=32)
    avatar = models.ImageField(verbose_name='头像', upload_to='images/avatar', default='images/avatar/default.jpg')
    create_time = models.DateTimeField(verbose_name='创建时间', auto_now_add=True)
    updated_time = models.DateTimeField(verbose_name='更新时间', auto_now=True)
    sign = models.CharField(verbose_name='个人签名', max_length=50,default=default_sign)
    info = models.CharField(verbose_name='个人简介', max_length=150, default='这个人很懒，还没有留下任何东西~~~')

    def __str__(self):
        return self.id

    def get_avatar_url(self):
        '''返回头像的url'''
        return MEDIR_ADDR + str(self.avatar)

    class Meta:
        db_table = 'user_User'

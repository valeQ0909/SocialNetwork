from django.http import JsonResponse
import tensorflow.keras as keras
import cv2
import numpy as np
import os

# InMemoryUploadedFile类的主要方法和属性包括：
#
# name: 上传文件的原始文件名。
# size: 上传文件的大小（以字节为单位）。
# content_type: 上传文件的MIME类型。
# read(): 读取文件内容并返回二进制数据。
# open(): 打开文件以进行读取或写入操作。
# close(): 关闭文件。
#
# 原文链接：https:// blog.csdn.net/samsion1/article/details/134001104

# DigitalMnist 手写数字识别模型
# https://www.jianshu.com/p/120f8099c924
def DigitalMnist(request):
    imgFile = request.FILES.get('image_file').read()  # imgFile是 'InMemoryUploadedFile' object
    #print("imgFile: " + str(imgFile))
    # opencv读取内存缓冲中的图片文件
    img = cv2.imdecode(np.fromstring(imgFile, np.uint8), cv2.IMREAD_COLOR)

    model = keras.models.load_model('E:\\GinkgoLab\\DeepLearning\\models\\ConvDigital_mnist.h5')
    # 注意该模型的输入是（x, 28, 28, 1），所以需要先reshape一下
    x_new = preHandle(img).reshape(1, 28, 28, 1)
    predict_x = model.predict(x_new)
    classes_x = np.argmax(predict_x, axis=1)

    err_msg = "predict: " + str(classes_x[0])
    print(err_msg)

    result = {'code': 10103, 'err_msg': err_msg, "predict": str(classes_x[0])}
    return JsonResponse(result)

############下面是自己的方法#############
def preHandle(img):
    img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)  # 3通道RGB转单通道灰度图
    img = (255 - np.array(img)) / 255.
    img = cv2.resize(img, (28, 28))
    return img

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
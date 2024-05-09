from PIL import Image

base_width = 300
img = Image.open('/home/jack/work/self/golang/gin-app/web/src-tauri/icons/head.png')
img_32_32 = img.resize((32,32))
img_128_128 = img.resize((128,128))
img_256_256 = img.resize((256,256))
img_32_32.save("/home/jack/work/self/golang/gin-app/web/src-tauri/icons/32x32.png")
img_128_128.save("/home/jack/work/self/golang/gin-app/web/src-tauri/icons/128x128.png")
img_256_256.save("/home/jack/work/self/golang/gin-app/web/src-tauri/icons/128x128@2x.png")

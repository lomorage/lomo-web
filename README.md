# lomo-web
Lomorage Web Application

Not a fully Web Lomorage client yet, currently it only supports uploading image and video files.

Download the binary and specify the address of lomo-backend service via `--baseurl` parameter, and then open browser to access lomo-web. For example if the Lomorage Service (lomod) is running at "192.168.1.12:8000", and you can either run lomo-web at different device with Lomorage Service or run lomo-web at the same device with Lomorage Service.

```
./lomo-web --baseurl http://192.168.1.12:8000
```

If running lomo-web at the same device with Lomorage Service, you can skip `--baseurl` and lomo-web will use the IP address on this device and default Lomorage service (lomod) port 8000.

```
./lomo-web
```

Then you can open broswer and access "http://[ip-of-lomo-web]", the default port is 80 and you can change that by using "--port" parameter when running "lomo-web".

[![Screen-Shot-2019-11-21-at-10-13-13-PM.png](https://i.postimg.cc/SNgbW2Kq/Screen-Shot-2019-11-21-at-10-13-13-PM.png)](https://postimg.cc/svGLz2S0)

[![Screen-Shot-2019-11-21-at-10-16-21-PM.png](https://i.postimg.cc/B64FbyG4/Screen-Shot-2019-11-21-at-10-16-21-PM.png)](https://postimg.cc/gwtjBgJT)
"use strict";
exports.__esModule = true;
var http = require("http");
var Spider = /** @class */ (function () {
    function Spider() {
        this.sum = 0;
    }
    Spider.prototype.httpGet = function (url) {
        return new Promise(function (resolve, reject) {
            http.get(url, function (res) {
                res.setEncoding('utf8');
                res.on('data', function (data) {
                    resolve(data);
                });
            }).on('error', function (e) {
                reject(e.message);
            });
        });
    };
    Spider.prototype.render = function (urls) {
        var _this = this;
        urls.reduce(function (promise, url) {
            return promise.then(function (_) {
                return _this.httpGet(url);
            }).then(function (node) {
                return _this.flush(node);
            }).then(function (res) {
                _this.sum += res;
            })["catch"](function (e) { return console.log(e); });
        }, Promise.resolve());
    };
    Spider.prototype.flush = function (data) {
        return new Promise(function (resolve) {
            var re = /:\s(\d+?);/gi;
            var number = Number(re.exec(String(data))[1]);
            resolve(number);
        });
    };
    return Spider;
}());
var baseUrl = 'http://dy-public.oss-cn-shenzhen.aliyuncs.com/interviewTestData/';
var urlList = [];
for (var index = 1; index <= 10; index++) {
    urlList.push("" + baseUrl + index + ".txt");
}
var spider = new Spider();
spider.render(urlList);
// 模拟等待微队列运行完成， 可以用await语法糖
// 这里只是写个大题的框架罢了，后面根据具体的要求改动
setTimeout(function () {
    console.log(spider.sum);
}, 5000);

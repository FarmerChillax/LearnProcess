"use strict";
exports.__esModule = true;
var http = require("http");
var url = "https://baidu.com";
http.get(url, function (res) {
    res.setEncoding('utf-8');
    console.log(res);
    res.on('data', function (data) {
        console.log(data);
    });
}).on('error', function (e) {
    console.log(e);
});

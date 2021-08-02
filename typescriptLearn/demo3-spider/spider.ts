import * as http from 'http'

class Spider {
    sum: Number
    constructor() {
        this.sum = 0
    }
    httpGet(url) {
        return new Promise((resolve, reject) => {
            http.get(url, res => {
                res.setEncoding('utf8')
                res.on('data', data => {
                    resolve(data)
                })
            }).on('error', (e) => {
                reject(e.message)
            })
        })
    }
    render(urls) {
        urls.reduce((promise, url) => {
            return promise.then(_ => {
                return this.httpGet(url)
            }).then(node => {
                return this.flush(node)
            }).then(res => {
                this.sum += res
            })
                .catch(e => console.log(e))
        }, Promise.resolve())
    }
    flush(data) {
        return new Promise(resolve => {
            const re = /:\s(\d+?);/gi
            const number = Number(re.exec(String(data))[1])
            resolve(number)
        })
    }
}

let baseUrl = 'http://dy-public.oss-cn-shenzhen.aliyuncs.com/interviewTestData/'
let urlList = []
for (let index = 1; index <= 10; index++) {
    urlList.push(`${baseUrl}${index}.txt`)
}

const spider = new Spider()
spider.render(urlList)
// 模拟等待微队列运行完成， 可以用await语法糖
// 这里只是写个大题的框架罢了，后面根据具体的要求改动
setTimeout(() => {
    console.log(spider.sum)
}, 5000);
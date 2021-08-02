import * as http from 'http'

let url = `https://baidu.com`

http.get(url, res => {
    res.setEncoding('utf-8')
    console.log(res)
    res.on('data', data => {
        console.log(data)
    })
}).on('error', e => {
    console.log(e)
})
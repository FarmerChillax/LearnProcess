const http = require('http');
const querystring = require('querystring')
let port = 5003

const users = [
    { "id": 1, 'name': 'farmer' },
    { "id": 2, 'name': 'xiaotao' }
]

function sleep(delay) {
    var start = (new Date()).getTime();
    while ((new Date()).getTime() - start < delay) {
        // 使用  continue 实现；
        continue; 
    }
}

const server = http.createServer((req, res) => {
    const method = req.method
    const url = req.url

    req.query = querystring.parse(url.split('?')[1])

    let name = "Not found"

    users.forEach(e => {
        if (e.name == req.query['name']) {
            name = e
        }
    })

    if (name == "Not found") {
        res.statusCode = 404
    }

    res.setHeader("Access-Control-Allow-Origin", "*")
    sleep(2000)
    res.end(JSON.stringify(name));
})

server.listen(port, () => {
    console.log('server running at port ', port);
})
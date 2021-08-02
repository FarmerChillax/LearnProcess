const http = require('http');
const querystring = require('querystring')
let port = 5003

const users = [
    { "id": 1, 'name': 'farmer' },
    { "id": 2, 'name': 'xiaotao' }
]

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

    res.setHeader("Access-Control-Allow-Origin", "*")
    res.end(JSON.stringify(name));
})

server.listen(port, () => {
    console.log('server running at port ', port);
})

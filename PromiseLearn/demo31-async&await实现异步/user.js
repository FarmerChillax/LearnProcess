const http = require('http');
const querystring = require('querystring')
let port = 5003

const users = [
    { "id": 1, 'name': 'farmer', "next": '2' },
    { "id": 2, 'name': 'xiaotao', "next": '1' }
]

const server = http.createServer((req, res) => {
    const method = req.method
    const url = req.url


    req.query = querystring.parse(url.split('?')[1])
    // console.log(req.query['name'])
    // console.log(url.split('?')[1])
    // console.log('**********************')
    let rep_body = "Not found"
    if (req.query['name']) {
        users.forEach(e => {
            if (e.name == req.query['name']) {
                rep_body = e
            }
        })
    }
    if(req.query['id']){
        users.forEach(e => {
            if (e.id == req.query['id']) {
                rep_body = e
            }
        })
    }

    // name += url
    res.setHeader("Access-Control-Allow-Origin", "*")
    res.end(JSON.stringify(rep_body));
})

server.listen(port, () => {
    console.log('server running at port ', port);
})

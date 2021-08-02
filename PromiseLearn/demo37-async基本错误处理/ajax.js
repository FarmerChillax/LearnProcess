class HttpError extends Error{
    constructor(msg){
        super(msg)
        this.name = 'HttpError'
    }
}

function ajax(url) {
    return new Promise((resolve, reject) => {
        let xhr = new XMLHttpRequest()
        xhr.open('GET', url)
        xhr.send()
        xhr.onload = function () {
            if (this.status == 200) {
                resolve(JSON.parse(this.response))
            } else {
                reject(new HttpError('用户不存在'))
            }
        }
        xhr.onerror = function () {
            reject(this)
        }
    })
}
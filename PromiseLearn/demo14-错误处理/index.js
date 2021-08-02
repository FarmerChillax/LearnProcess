// 栗子一
new Promise((resolve, reject) => {
    reject(new Error("Promise error message."))
}).then(
    value => {
        console.log(value)
    }, reason => {
        console.log(reason)
        console.log(reason.message)
    }
)
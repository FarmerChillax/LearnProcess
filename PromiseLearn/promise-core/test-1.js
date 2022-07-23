let testPromise = new PromiseCore((resolve, reject) => {
    resolve("test-1")
}).then(res => {
    res("1")
})
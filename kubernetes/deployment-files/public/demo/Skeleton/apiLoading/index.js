document.onreadystatechange = () => {
    const Skeleton = document.querySelector("#Skeleton")
    Skeleton.classList.add("loading")

    const text = document.querySelectorAll(".content h4")
    text.forEach(e => { e.innerText = "" })

    const description = document.querySelectorAll(".description")
    description.forEach(e => { e.innerText = "" })

    if (document.readyState == "complete") {
        console.log(Skeleton.classList.remove("loading"))
        text.forEach(e => { e.innerText = "Skeleton Demo" })
        description.forEach(e => {
            e.innerText = "骨架框效果测试样例，这是一段测试使用噶文字。This is a test text."
        })
    }
}


// test code

// 判断图片是否加载完
function imgLoad(img, callback) {
    const timer = setInterval(function() {
        if(img.complete) {
            callback(img)
            clearInterval(timer)
        }
    }, 500)
}

const img = document.querySelector(".image img")
imgLoad(img,function() {
    console.log("loaded");
})

// function changeValue(node, value) {
//     node.innerText = value
// }
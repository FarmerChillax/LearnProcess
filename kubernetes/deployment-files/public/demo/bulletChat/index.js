const btn = document.querySelector("button")
const input = document.querySelector("input")
// 模拟发送阻塞时间
let timeout = 3
// 弹幕列表
let bulletList = []
let bulletY = [18, 36, 54]

// click
function click() {
    const val = input.value
    if (val.trim() != "") {
        input.value = ""
        this.className = "btn_active"
        this.innerHTML = `${timeout}s`
        // 模拟网络IO等待时间
        const netWorkIO = setInterval(() => {
            timeout--
            this.innerHTML = `${timeout}s`
            if (timeout === 0) {
                clearInterval(netWorkIO)
                this.className = ""
                this.innerHTML = `发送`
                timeout = 3
            }
        }, 1000)
        // 添加弹幕内容到弹幕列表
        bulletList.push({
            text: val,
            x: 600,
            y: bulletY[parseInt(Math.random() * 3)] // y轴随机
        })
    }
}

btn.addEventListener("click", click)
input.addEventListener("keyup", (e) => {
    if (e.keyCode  === 13 && btn.className != "btn_active") {
        btn.click()
    }
})

// canvas部分
const canvas = document.querySelector("canvas")
// open 2D draw
const ctx = canvas.getContext("2d")

ctx.font = `${bulletY[0]}px 微软雅黑`

ctx.fillStyle = "#fff"
// 每帧不同位置（移动）
const draw = () => {
    // 从右向左移动（ 600 减到 0），每秒60帧
    bulletList.forEach(item => {
        item.x--
    })
    // 获取电脑屏幕刷新率
    requestAnimationFrame(() => {
        ctx.clearRect(0, 0, 600, 250)
        // 绘制弹幕
        bulletList.forEach(item => {
            ctx.fillText(item.text, item.x, item.y)
        })
    })
    setTimeout(() => {
        draw()
    }, 1000 / 60);
}

draw()
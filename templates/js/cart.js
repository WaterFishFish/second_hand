const token = window.localStorage.getItem('token')
const id = window.localStorage.getItem('id')

if (!token || !id) {
    window.location.href ='login'
}
else {
    getcartList()
}
function getcartList(){
    $.ajax({
        url: 'http://localhost:8080/cart/list',
        method: 'Get',
        data: { id: id },
        headers: { authorization: token },
        success(res) {
            if (res.code !== 1) {
                window.location.href = '/login'
                return
            }

            bindHtml()
        }
    })
}

function bindHtml(res) {
    if (res.cart.length) {
        if (!res.cart.length) {
            $('.empty').addClass('active')
            $('.list').removeClass('active')
            return
        }
    }
    
    console.log(res.cart)

       let selectNum = 0, totalPrice = 0, totalNum = 0
    res.cart.forEach(item => {
        if (item.is_select) {
            selectNum++
            totalNum += item.cart_number;
            total_price = item.cart_number * item.price
        }
    });

        let str = `
    <div class="top">
              全选  <input type="checkbox" ${ selectNum === res.cart.length ? 'checked' : ''}>
            </div>
            <ul class="center">
    `

res.cart.forEach(item => {
    str += `
    <li>
           <div class="select">
                        <input type="checkbox" bookId="${item.id}">
                    </div>
                    <div class="show">
                        <img src="${item.img_small_logo}" alt="">
                    </div>
                    <div class="title">
                        ${item.title}
                    </div>
                    <div class="price">${item.price}</div>
                    <div class="number">
                        <button>-</button>
                        <input type="text" value=${item.cart_number}>
                        <button>+</button>
                    </div>
                    <div class="subPrice">￥ ${(item.price * item.cart_number).roFixed(2)}</div>
                    <div class="destory">
                        <button></button>
                    </div>
        `
    });

    str += `
                </li>
            </ul>
            <div class="bottom">
                <p>
                    共计 <span>${totalNum}</span> 件商品
                </p>
                <div class="btns">
                    <button>清空购物车</button>
                    <button ${selectNum === 0 ? 'disabled' : ''}>删除所有已选中</button>
                    <button ${selectNum === 0 ? 'disabled' : ''}>去支付</button>
                </div>
                <p>
                    共计 <span>${totalPrice}</span>
                </p>
            </div>
     `
    
    $('.list').html(str)
}

$('.list').on('.click', '.center .select input', function () {
    console.log('修改选中状态')

})
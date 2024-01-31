// const id = window.localStorage.getItem('bookId')
const id = window.localStorage.getItem('bookId');

const info ={
    id:id,
    token:''
}
info.token=window.localStorage.getItem('token');
console.log(info.token)




GetBookInfo()
function GetBookInfo() {

    // $.ajax({
    //     url: 'http://localhost:8080/home/DirectPurchase',
    //     method: 'GET',
    //     data: { id: id, token: token },
    //     success: function(res) {
    //         // 在成功回调中渲染页面或处理响应数据
    //         console.log(res);
    //         renderProductPage(res);
    //     },
    //     error: function(err) {
    //         // 处理错误情况
    //         console.error('Error:', err);
    //     }
    // });

    $.get('http://localhost:8080/home/DirectPurchase', info,res=> {
        // 在成功回调中渲染页面
        console.log(res)
        renderProductPage(res);
    });
}




        // 渲染页面的函数
        function renderProductPage(productData) {
            // 获取到的产品数据可以在这里进行处理
            // 例如，你可以通过 productData.image 获取商品图片的 URL
            // 构建HTML结构
            productData.src ='../../imgs/' + productData.src
            let html = `
<div class="container">
                <div id="product-image">
                    <img src="${productData.src}" alt="商品图片" >
                </div>
                <div id="product-info">
                    <span>${productData.bookName}</span>
                    <p>${productData.description}</p>
                    <p>价格: ￥${productData.price}</p>
                </div>
                <button id="buy-button">购买</button>
                <div id="user-balance">
                    用户余额: ￥${productData.Balance}
                </div>
                <div id="success-message"></div>
                
                    <div id="success-message"></div>
</div>
            `;


            // 将HTML插入到容器中
            $('.list').html(html);

// 监听购买按钮的点击事件等其他操作
            $('#buy-button').on('click', function () {
                // 这里可以添加购买逻辑，例如向服务器发送购买请求
                let url = "/home/DirectPurchase/" + id;
                // $.post(url, info)

                $.ajax({
                    url: url,
                    method: 'GET',
                    data: { id: info.id, token: info.token },
                    success: function(res) {
                        // 在成功回调中渲染页面或处理响应数据
                        console.log(res);
                      //  renderProductPage(res);
                    },
                    error: function(err) {
                        // 处理错误情况
                        console.error('Error:', err);
                    }
                });

                // 成功购买后，显示成功消息
                $('#success-message').text('购买成功！');
                alert('购买成功');
                window.location.href='/home/list';
            });


        }


const bookId = window.localStorage.getItem('bookId')
console.log(bookId)
if (!bookId) window.location.href = '/home/bookList'
const info ={
    id: bookId,
    test: 1
}

GetBookInfo()
function GetBookInfo() {
    $.get('http://localhost:8080/home/bookList/detail', info, res => {
        // console.log(res)
        res.src += '../../../imgs/'+res.src
        $('.show > img').prop('src', res.src)
        $('.info > .title').text(res.description)
        $('.info > .price').text('￥'+res.price)
    })
}


$('.buy').on('click', function () {
    // 在这里设置跳转的页面地址
    window.location.href = '/home/DP'; // 替换为实际的页面地址
});
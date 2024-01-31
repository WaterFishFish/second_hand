// 调用函数发送请求

let token=window.localStorage.getItem('token');
getcatList()
function getcatList()
{
    $.get('http://localhost:8080/home/getList', res => {
        console.log(res)

        let str = '<li class="active">全部</li>'
        res.list.forEach( item => {
            str += `<li>${ item }</li>`
        })
        $('.category').html(str)
    })

}




const info ={
    current: 1,
    pagesize:12,
    search: '',
    total: 1,
    class:'全部',
    time:new Date().toISOString(),
    sortType:'',
    AD:'',
    token: ''
}
info.token=window.localStorage.getItem('token')
let totalPage=1
console.log(info)
function getGoodsList() {

    $.get('http://localhost:8080/home/bookList',info,res => {
        console.log(res.pagesize)
        info.current=res.current;
        info.total=res.total;
        info.sortType=res.sortType;
        info.class=res.class;
        info.pagesize=res.pagesize;
        info.AD=res.AD;
        console.log(res.pagesize)
        bindHtml(res);
    })

}
getGoodsList()

function bindHtml(res) {
    // console.log(res)

    if (info.current === 1) $('.left').addClass('disable')
    else $('.left').removeClass('disable')

    if (info.current === res.total) $('.right').addClass('disable')
    else $('.right').removeClass('disable')
    $('.total').text(`${ info.current} / ${ res.total }`)

    $('select').val(info.pagesize)

    $('.page').val(info.current)

    let str = ``
    res.list.forEach(item => {
        // console.log(item.src)
        str += `
    <li goodsNumber="${ item.ID}">
        <div class="show">
            <img src="../../imgs/${ item.src }" alt="">
        </div>
        <div class="info">
            <p class="title">${ item.description}</p>
            <p class="price">
                <span class="current">${'￥' + item.price }</span>
                <span class="old"></span>
            </p>
            <button bookId=${item.ID}>加入购物车</button>
        </div>
    </li>
`
    })
    $('.list').html(str)

}

$('.category').on('click', 'li', function () {
    $(this).addClass('active').siblings().removeClass('active')

    info.category = $(this).text === '全部' ? '' : $(this).text()
    info.current = 1
    info.class = $(this).text()
    getGoodsList()
})

// $('.filter').on('click', 'li', function () {
//     $(this).addClass('active').siblings().removeClass('active')

// })

$('.sort').on('click', 'li', function (){
    $(this).addClass('active').siblings().removeClass('active')

    info.sortType = $(this).attr('type')
    info.sortMethod = $(this).attr('method')
    // info.sortType =
    getGoodsList()
})

$('search').on('input', function () {

    info.search = $(this).val().trim()
    info.current = 1

    getGoodsList()
})

$('.left').on('click', function () {
    if ($(this).hasClass('disable')) return

    info.current--
    console.log(info.current)

    getGoodsList()
})

$('.right').on('click', function () {
    if ($(this).hasClass('disable')) return

    info.current++
    console.log(info.current)
    getGoodsList()
})

$('select').on('change', function () {

    info.pagesize = $(this).val()
    info.current = 1

    getGoodsList()
})

$('.jump').on('click', function () {
    let page = $('.page').val()

    if (isNaN(page)) page = 1

    if (page <= 1) page = 1

    if (page >= totalPage) page = totalPage

    info.current = page

    getGoodsList()
})

// 加入购物车
$('.list').on('click', 'button', function(e){
    e.stopPropagation()

    const token = window.localStorage.getItem('token')
    const sid = $(this).attr('bookid')
    $.ajax({
        url: 'http://localhost:8080/home/ShoppingCarts/add',
        method: 'POST',
        headers: { authorization: token },
        data: { bookId: sid },
        success() {
            console.log(sid)
            window.alert('加入购物车成功')
            return
        }
    })
})

$('.list').on('click', 'li', function () {
    console.log('跳转到详情页面')

    window.localStorage.setItem('bookId', $(this).attr('goodsnumber'))
    console.log(window.localStorage.getItem('bookId'))
window.localStorage.getItem('bookId')
    window.location.href = '/home/detail'

})
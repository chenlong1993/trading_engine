(function(){
    layui.use(['layer'], function () {
        var layer = layui.layer;

        var login = {
            init: function(){
                if(!Cookies.get("user_id")) {
                    layer.prompt({
                        formType: 0,
                        value: "1000",
                        title: '第一次，请设置您的用户ID',
                        area: ['400px', '50px'] // 自定义文本域宽高
                    }, function(value, index, elem){
                        Cookies.set("user_id", value, { expires: 7, path: '' });
                        layer.close(index); // 关闭层
                    });
                }else{
                    layui.$(".user").html(Cookies.get("user_id"));
                    
                }
            }
        };
        login.init();
    })

})()
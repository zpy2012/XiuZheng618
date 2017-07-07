//style.html
$("tr .detail_edit_button").click(function(){

	var list = $(this).parents(".zjy_style_list tr").index();
	var URL = "/admin/detail";
	//这里取了id
 	var id= $(".DeatilId").eq(list).text();
  var that = $(this);
	$.post(URL,
    {
      id:id
    },
    function(data,status){
      if (data == "成功") {
      	 that.html("已领奖");
         that.attr("disabled","disabled");
      }
    });
});
 
//chart.html
var arrays = [];
//从页面上隐藏
$(".zjy_lists .am-btn-danger").click(function(){
	var list = $(this).parents(".zjy_lists").index();
//  		$(".zjy_lists").eq(list).remove();
	$(".zjy_lists").eq(list).css("display","none");
	if($(".zjy_lists input[type='checkbox']").eq(list).is(':checked')){ 
		arrays.splice($.inArray(list,arrays),1);
		console.log(arrays)
	}
});
//复选框选中的是列表中的第几个
$(".zjy_lists input[type='checkbox']").click(function(){
	var list = $(this).parents(".zjy_lists").index();
	if($(this).is(':checked')){ 
		arrays.push(list);
		if(arrays.length>3){
			var shiftfirst = arrays.shift();
			$(".zjy_lists input[type='checkbox']").eq(shiftfirst).uCheck('uncheck');
		}
	}else{
		arrays.splice($.inArray(list,arrays),1)
	}
	console.log(arrays)
});

//index.html
//从页面上删除
$(".zjy_list .index_delete_button").click(function(){
	var URL = "/admin/index/delete";
  	var id = $("#indexId").text();
	$.post(URL,
    {
      id:id
    },
    function(data,status){
      if (data == "成功") {
      	window.location.reload();
      }
    });
});

//chart.html
//从页面上删除
$(".zjy_listbox .chart_delete_button").click(function(){
	var URL = "/admin/chart/delete";
	//这里取了点击是第几个
	var list = $(this).parents(".zjy_listbox").index();
	//这里取了id
 	var id= $(".ChartId").eq(list).text();
	$.post(URL,
    {
      id:id
    },
    function(data,status){
      if (data == "成功") {
      	window.location.reload();
      }
    });
});

//单选框选中的是列表中的第几个
$(".zjy_list input[type='radio']").click(function(){
	var list = $(this).parents(".zjy_list").index();
});


//index.html
//编辑
$(".zjy_listbox .index_edit_button").click(function(){
	var URL = "/admin/index/edit";
  	var id = $("#indexId").text();
	$.post(URL,
    {
      id:id,
      edit:true
    },
    function(data,status){
      if (data == "成功") {
      	window.location.href=URL;
      }
    });
});

//chart.html
//编辑
$(".zjy_listbox .chart_edit_button").click(function(){
	var URL = "/admin/chart/edit";
	//这里取了点击是第几个
	var list = $(this).parents(".zjy_listbox").index();
	//这里取了id
 	var id= $(".ChartId").eq(list).text();
	$.post(URL,
    {
      id:id,
      edit:true
    },
    function(data,status){
      if (data == "成功") {
      	window.location.href=URL;
      }
    });
});

//index新增
$(".index_new").click(function(){
	window.location.href="/admin/index/add";
});

//chart新增
$(".chart_new").click(function(){
	window.location.href="/admin/chart/add";
});

//style.html
$("tr .style_edit_button").click(function(){
	var list = $(this).parents(".zjy_style_list tr").index();
	var URL = "/admin/style/edit";
	//这里取了id
 	var id= $(".StyleId").eq(list).text();
	$.post(URL,
    {
      id:id,
      edit:true
    },
    function(data,status){
      if (data == "成功") {
      	window.location.href=URL;
      }
    });
});
var stylearrays = [];
$(".style_delete_button").click(function(){
	var list = $(this).parents(".zjy_style_list tr").index();
	var URL = "/admin/style/delete";
	//这里取了id
 	var id= $(".StyleId").eq(list).text();
	$.post(URL,
    {
      id:id
    },
    function(data,status){
      if (data == "成功") {
      	window.location.reload();
      }
    });
});
$(".zjy_style_list input[type='checkbox']").click(function(){
	var list = $(this).parents(".zjy_style_list tr").index();
	if($(this).is(':checked')){ 
		stylearrays.push(list);
		if(stylearrays.length>3){
			var shiftfirst = stylearrays.shift();
			$(".zjy_style_list input[type='checkbox']").eq(shiftfirst).uCheck('uncheck');
		}
	}else{
		stylearrays.splice($.inArray(list,stylearrays),1)
	}
	console.log(stylearrays)
});
$(".zjy_style_new").click(function(){
	window.location.href="/admin/style/add";
});
//users.html
$("tr .see_userinfo_button").click(function(){
	var list = $(this).parents(".zjy_user_list tr").index();
	var id = $(".UserId").eq(list).text();
  var URL = "/admin/userlist/useroperation/?page=1&id="+id;
  window.location.href=URL;
});


//product.html
//新增产品
$(".zjy_product_new").click(function(){
	window.location.href="/admin/products/add";
});
$(".zjy_product_list .am-btn-secondary").click(function(){
	var URL = "/admin/products/edit";
	//这里取了点击是第几个
	var list = $(this).parents(".zjy_product_list").index();
	//这里取了id
 	var id= $(".FlowerId").eq(list).text();
	$.post(URL,
    {
      id:id,
      edit:true
    },
    function(data,status){
      if (data == "成功") {
      	window.location.href=URL;
      }
    });
});
$(".zjy_product_list .product_delete_button").click(function(){
	var list = $(this).parents(".zjy_product_list").index();//点击的第几个编辑
	var URL = "/admin/products/delete";
	//这里取了id
 	var id= $(".FlowerId").eq(list).text();
	$.post(URL,
    {
      id:id
    },
    function(data,status){
      if (data == "成功") {
      	window.location.reload();
      }
    });
});
//imgform-lines.html
//增加图片
$("#doc-form-file").change(function() {
var $file = $(this);
var fileObj = $file[0];
var windowURL = window.URL || window.webkitURL;
var dataURL;
var $img = $("#preview");
if(fileObj && fileObj.files && fileObj.files[0]){
dataURL = windowURL.createObjectURL(fileObj.files[0]);
$img.attr('src',dataURL);
}else{
dataURL = $file.val();
var imgObj = document.getElementById("preview");
imgObj.style.filter = "progid:DXImageTransform.Microsoft.AlphaImageLoader(sizingMethod=scale)";
imgObj.filters.item("DXImageTransform.Microsoft.AlphaImageLoader").src = dataURL;
 
}
});

//处理form表单的返回数据
// $(document).on("submit", ".ajax-form", function (e) {
//     var submitButton = $(this).find("button[type=submit]");
//     submitButton.attr("disabled", "disabled").text("正在提交");
//     e.preventDefault();
//     var action = $(this).attr("action");
//     var method = $(this).attr("method");
//     $.ajax({
//         url: action,
//         method: method,
//         type:"post",
//     	processData: false,
//     	contentType: false,
//         // data: $(this).serialize(),
//         success: function (data) {
        	
       //      if (data.Code == undefined) {
       //      	$(".am-modal-bd").text(data);
		     //  	$('#your-modal').modal('open');
			  		// setTimeout(function(){
			  		// 	$('#your-modal').modal('close')
			  		// },10000)//2秒后自动关闭
       //          submitButton.removeAttr("disabled").text("提交");
       //          return;
       //      }
       //      if (data.Code == 200) {
       //          setTimeout(function () {
       //              if (data.JumpUrl != "") {
       //                  window.location.href = data.JumpUrl;
       //              }
       //              else {
       //                  location.reload();
       //              }
       //          }, 2500);
       //      }
//         },
//         error: function () {
//             submitButton.removeAttr("disabled").text("提交");
//         }
//     })
// });
$(".submit_button").click(function() {
    $(".am-modal-bd").text("正在上传，请稍后....");
    $('#your-modal').modal('open');
    var option = {
        headers : {
            "ClientCallMode" : "ajax"
        }, //添加请求头部
        dataType: "json",
        success : function(data) {
            $('#your-modal').modal('close');
            if (data.Code == "200") {
                $(".am-modal-bd").text(data.Msg);
  		      	  $('#your-modal').modal('open');
      		  		setTimeout(function(){
      		  			$('#your-modal').modal('close')
      		  		},1000);//2秒后自动关闭
                setTimeout(function () {
                    window.location.href = data.JumpUrl;
                }, 1500);
                return false
            }else {
                $(".am-modal-bd").text(data.Msg);
    		      	$('#your-modal').modal('open');
    			  		setTimeout(function(){
    			  			$('#your-modal').modal('close')
    			  		},1000);//2秒后自动关闭
                    return false;
                }
        },
        error : function(data) {
        	$(".am-modal-bd").text(data.Msg);
	      	$('#your-modal').modal('open');
		  		setTimeout(function(){
		  			$('#your-modal').modal('close')
		  		},10000);//2秒后自动关闭
            return false;
        }
    };
    $(".ajax-form").ajaxSubmit(option);
    return false; //最好返回false，因为如果按钮类型是submit,则表单自己又会提交一次;返回false阻止表单再次提交      
});